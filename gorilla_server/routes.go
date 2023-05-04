package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s Server) routes() http.Handler {
	router := mux.NewRouter()

	publicRoutes := router.Methods(http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete).Subrouter()
	pvtRoutes := router.Methods(http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete).Subrouter()

	publicRoutes.Methods(http.MethodGet).Path("/pokemon").HandlerFunc(HandlePokemon)
	publicRoutes.Methods(http.MethodPost).Path("/v1/config/cluster").HandlerFunc(LogRequest)
	publicRoutes.Methods(http.MethodGet).Path("/v1/config/cluster/load-env").HandlerFunc(LogRequest)

	// cloud inventory routes
	pvtRoutes.Methods(http.MethodGet).Path("/v1/inventory/aws/global").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodGet).Path("/v1/inventory/aws/projects/{projectId}/environments/{envId}/account").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodGet).Path("/v1/inventory/azure/global").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodGet).Path("/v1/inventory/azure/projects/{projectId}/environments/{envId}/account").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodPost).Path("/v1/inventory/aws/global/resync").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodPost).Path("/v1/inventory/aws/projects/{projectId}/account/re-sync").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodPost).Path("/v1/inventory/azure/global/resync").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodPost).Path("/v1/inventory/azure/projects/{projectId}/account/re-sync").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodGet).Path("/v1/inventory/gcp/global").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodGet).Path("/v1/inventory/gcp/projects/{projectId}/environments/{envId}/account").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodPost).Path("/v1/inventory/gcp/global/resync").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodPost).Path("/v1/inventory/gcp/projects/{projectId}/account/re-sync").HandlerFunc(LogRequest)

	// environment config routes
	pvtRoutes.Methods(http.MethodGet).Path("/v1/projects/{projectId}/environments/list").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodPost).Path("/v1/projects/{projectId}/environments/create").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodPut).Path("/v1/projects/{projectId}/environments/{envId}/update").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodPut).Path("/v1/projects/{projectId}/environments/{envId}/terraform/unset").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodDelete).Path("/v1/projects/{projectId}/environments/{envId}/delete").HandlerFunc(LogRequest)
	pvtRoutes.Methods(http.MethodGet).Path("/v1/projects/{projectId}/environments/{envId}/terraform").HandlerFunc(LogRequest)

	pvtRoutes.Use(SleeperMiddleware(time.Second * 3))

	return router
}
