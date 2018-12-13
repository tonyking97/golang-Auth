package main

import (
	"./routers"
	"net/http"
)

func main(){
	router := routers.InitRouters()
	http.ListenAndServe(":8080",router)
}

//TODO common configuration file should be implemented
//TODO Init folder should be implemented and necessary code should move into that