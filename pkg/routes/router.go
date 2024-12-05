package router

import (
	"net/http"

	v1router "github.com/api-assignment/pkg/routes/v1"
	"github.com/go-chi/chi/v5"
)

func MainRouter() http.Handler {
	r := chi.NewRouter()
	r.Mount("/api", registerRouterVersions())
	return r
}

func registerRouterVersions() http.Handler {
	r := chi.NewRouter()
	r.Mount("/v1", v1router.V1Router())
	return r

}
