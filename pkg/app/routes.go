package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV2 = "/api/v2"

	// specification routes
	a.Router.HandleFunc(apiV2+"/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/specification/test", a.testConnectorSpecification).Methods("POST", "OPTIONS")

	// tickets routes
	a.Router.HandleFunc(apiV2+"/tickets", a.getAllTickets).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/tickets/metadata", a.getAllTicketsMetadata).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/tickets/{id}", a.getSingleTicket).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/tickets", a.createTicket).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/tickets/{id}", a.deleteTicket).Methods("DELETE", "OPTIONS")

	// search
	a.Router.HandleFunc(apiV2+"/search", a.searchTickets).Methods("GET", "OPTIONS")

	// stats
	a.Router.HandleFunc(apiV2+"/stats/agent/{agent_id}", a.getStatsForAgent).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/stats/group/{group_id}", a.getStatsForGroup).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/stats/date/{created_at}", a.getStatsForTicketsSinceDate).Methods("GET", "OPTIONS")

	// agents
	a.Router.HandleFunc(apiV2+"/agents", a.getAgents).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/agents/{id}", a.getSingleAgent).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/agents/{id}", a.deactivateSingleAgent).Methods("DELETE", "OPTIONS")

	// problems
	a.Router.HandleFunc(apiV2+"/problems", a.getProblems).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/problems/{id}", a.getSingleProblem).Methods("GET", "OPTIONS")

	// groups
	a.Router.HandleFunc(apiV2+"/groups", a.getGroups).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/groups/{id}", a.getSingleGroup).Methods("GET", "OPTIONS")

	// assets
	a.Router.HandleFunc(apiV2+"/assets", a.getAssets).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/assets/{id}", a.getSingleAsset).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV2+"/assets/{id}", a.deleteSingleAsset).Methods("DELETE", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
