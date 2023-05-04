package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func (s Server) routes() http.Handler {
	r := chi.NewMux()
	r.Route("/v1", func(r chi.Router) {
		r.Group(registerPublicRoutes)
		r.Group(registerPrivateRoutes)
	})
	return r
}

func registerPublicRoutes(r chi.Router) {
	r.Route("/config/cluster", func(r chi.Router) {
		r.Post("/", LogRequest)
		r.Get("/load-env", LogRequest)
	})
}

func registerPrivateRoutes(r chi.Router) {
	r.Use(SleeperMiddleware(time.Second * 3))

	// cloud inventory routes
	r.Route("/inventory", registerInventoryRoutes)

	// environment config routes
	r.Route("/projects/{projectId}/environments", registerProjectEnvRoutes)
}

func registerProjectEnvRoutes(r chi.Router) {
	r.Get("/list", LogRequest)
	r.Post("/create", LogRequest)
	r.Put("/{envId}/update", LogRequest)
	r.Put("/{envId}/terraform/unset", LogRequest)
	r.Delete("/{envId}/delete", LogRequest)
	r.Get("/{envId}/terraform", LogRequest)
}

func registerInventoryRoutes(r chi.Router) {
	// aws routes:
	r.Route("/aws", func(r chi.Router) {
		r.Get("/global", LogRequest)
		r.Post("/global/resync", LogRequest)
		r.Get("/projects/{projectId}/environments/{envId}/account", LogRequest)
		r.Post("/projects/{projectId}/account/re-sync", LogRequest)
	})

	// azure routes:
	r.Route("/azure", func(r chi.Router) {
		r.Get("/global", LogRequest)
		r.Post("/global/resync", LogRequest)
		r.Get("/projects/{projectId}/environments/{envId}/account", LogRequest)
		r.Post("/projects/{projectId}/account/re-sync", LogRequest)
	})

	// gcp routes:
	r.Route("/gcp", func(r chi.Router) {
		r.Get("/global", LogRequest)
		r.Post("/global/resync", LogRequest)
		r.Get("/projects/{projectId}/environments/{envId}/account", LogRequest)
		r.Post("/projects/{projectId}/account/re-sync", LogRequest)
	})
}
