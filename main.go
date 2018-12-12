package main

import (
	"./routers"
	"net/http"
)

func main(){
	router := routers.InitRouters()
	http.ListenAndServe(":8080",router)
}
