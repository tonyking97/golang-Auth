package routers

import (
	"../core"
	"../controllers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func SetAuthenticationRouter(router *mux.Router){
	commonHandler := alice.New(core.LoggingMiddleware)
	router.Handle("/",commonHandler.ThenFunc(controllers.LoginHandler)).Methods("POST")
	router.Handle("/signup",commonHandler.ThenFunc(controllers.SignUpHandler)).Methods("POST")
	//TODO Add Handler to check token is valid
}

