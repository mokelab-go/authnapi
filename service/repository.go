package service

import "github.com/mokelab-go/authnapi"

// AccountRepository provides account related API
type AccountRepository interface {
	// GetWithIdentifierAndPassword gets account with id and password
	GetWithIdentifierAndPassword(identifier, password string) (Account, error)

	// CreateSession creates new session
	CreateSession(accountID string) (authnapi.Session, error)

	// CreateRefresh creates new refresh token
	CreateRefresh(accountID string) (authnapi.Session, error)

	// GetRefresh gets refresh token
	GetRefresh(token string) (authnapi.Session, error)
}

// Account is app Account
type Account interface {
	ID() string
	ToJSON() map[string]interface{}
}
