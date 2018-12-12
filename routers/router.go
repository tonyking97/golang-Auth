package routers

import (
	"github.com/gorilla/mux"
)

func InitRouters() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	authenticationRouter := router.PathPrefix("/").Subrouter().StrictSlash(true)

	SetApiRouter(apiRouter)
	SetAuthenticationRouter(authenticationRouter)

	return router
}
