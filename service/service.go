package service

import (
	"errors"
	"net/http"

	"github.com/mokelab-go/authnapi"
	"github.com/mokelab-go/server/entity"
)

var (
	errEmptyIdentifier = errors.New("empty identifier")
	errEmptyPassword   = errors.New("empty password")
	errEmptyToken      = errors.New("empty token")
	errWrongIDPass     = errors.New("wrong identifier or password")
)

type service struct {
	accountRepo AccountRepository
}

// New creates service instance
func New(accountRepo AccountRepository) authnapi.Service {
	return &service{
		accountRepo: accountRepo,
	}
}

func (s *service) Authenticate(identifier, password string) entity.Response {
	if len(identifier) == 0 {
		return errorResponse(http.StatusBadRequest, "input_error", errEmptyIdentifier)
	}
	if len(password) == 0 {
		return errorResponse(http.StatusBadRequest, "input_error", errEmptyPassword)
	}

	a, err := s.accountRepo.GetWithIdentifierAndPassword(identifier, password)
	if err != nil {
		return errorResponse(http.StatusBadRequest, "input_error", errWrongIDPass)
	}

	accountID := a.ID()
	session, err := s.accountRepo.CreateSession(accountID)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, "server_error", err)
	}

	refresh, err := s.accountRepo.CreateRefresh(accountID)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, "server_error", err)
	}

	return entity.Response{
		Status: http.StatusOK,
		Body: map[string]interface{}{
			"account_id":    accountID,
			"access_token":  session.ID(),
			"refresh_token": refresh.ID(),
		},
	}
}

func (s *service) Refresh(token string) entity.Response {
	if len(token) == 0 {
		return errorResponse(http.StatusBadRequest, "input_error", errEmptyToken)
	}

	refresh, err := s.accountRepo.GetRefresh(token)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, "server_error", err)
	}

	accountID := refresh.AccountID()
	session, err := s.accountRepo.CreateSession(accountID)
	if err != nil {
		return errorResponse(http.StatusInternalServerError, "server_error", err)
	}

	return entity.Response{
		Status: http.StatusOK,
		Body: map[string]interface{}{
			"account_id":   accountID,
			"access_token": session.ID(),
		},
	}
}

func errorResponse(status int, code string, err error) entity.Response {
	return entity.Response{
		Status: status,
		Body: map[string]interface{}{
			"code": code,
			"msg":  err.Error(),
		},
	}
}
