package v1router

import (
	"net/http"

	"github.com/api-assignment/pkg/contoller"
	authMiddleware "github.com/api-assignment/pkg/middleware/auth"
	"github.com/go-chi/chi/v5"
)

func V1Router() http.Handler {
	r := chi.NewRouter()
	r.Mount("/auth", authRouter())
	r.Mount("/", protectedRouter())
	return r
}

func authRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/signup", contoller.Signup)
	r.Post("/login", contoller.Login)
	r.Get("/token", contoller.RefreshAccessToken)
	return r
}

func protectedRouter() http.Handler {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Use(authMiddleware.AccessTokenVerify)
		r.Get("/me", contoller.CheckIfSessionValid)
		r.Get("/user", contoller.GetUserData)
		r.Patch("/deactivate", contoller.DeActivateUser)
		r.Patch("/changePassword", contoller.ChangePassword)
	})
	return r
}
