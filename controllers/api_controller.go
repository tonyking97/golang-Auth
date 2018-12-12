package controllers

import (
	"fmt"
	"net/http"
)

func ApiHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hello API")
}
