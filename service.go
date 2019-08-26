package authnapi

import (
	"github.com/mokelab-go/server/entity"
)

// Service is service
type Service interface {
	// Authenticate is called in POST /auth
	Authenticate(identifier, password string) entity.Response
	// Refresh is called in POST /auth/refresh
	Refresh(token string) entity.Response
}
