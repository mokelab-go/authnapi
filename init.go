package authnapi

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mokelab-go/hop"
)

// InitRouter is an entry point
func InitRouter(router *mux.Router, s Service) {
	// authenticate (log in)
	router.Methods(http.MethodPost).
		Path("/auth").
		Handler(hop.Operations(
			hop.GetBodyAsJSON,
		)(authenticate(s)))

	// refresh token
	router.Methods(http.MethodPost).
		Path("/auth/refresh").
		Handler(hop.Operations(
			hop.GetBodyAsJSON,
		)(refresh(s)))

}
