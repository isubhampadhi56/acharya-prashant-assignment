package v1router

import (
	"net/http"

	"github.com/api-assignment/pkg/contoller"
	"github.com/go-chi/chi/v5"
)

func V1Router() http.Handler {
	r := chi.NewRouter()
	r.Mount("/auth", authRouter())
	return r
}

func authRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/signup", contoller.Signup)
	r.Post("/login", contoller.Login)
	return r
}
