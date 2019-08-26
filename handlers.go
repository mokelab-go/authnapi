package authnapi

import (
	"net/http"

	"github.com/mokelab-go/hop"
)

func authenticate(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := hop.BodyJSON(r.Context())
		identifier, _ := params["identifier"].(string)
		password, _ := params["password"].(string)

		resp := s.Authenticate(identifier, password)
		resp.Write(w)
	}
}

func refresh(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := hop.BodyJSON(r.Context())
		token, _ := params["token"].(string)

		resp := s.Refresh(token)
		resp.Write(w)
	}
}
