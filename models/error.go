package models

import "net/http"

type ErrorMessage struct {
	Message string `json:"err"`
}

func CheckError(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}