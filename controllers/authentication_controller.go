package controllers

import (
	"../core"
	"../models"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	models.CheckError(err,w)

	//Validating params. Refer models/users.go
	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &models.ErrorMessage{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err,w)

	} else {
		responseStatus, token := core.Login(requestUser)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseStatus)
		_, err := w.Write(token)
		models.CheckError(err,w)
	}
}

//TODO RefreshTokenHandler need to optimized
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	models.CheckError(err,w)

	w.Header().Set("Content-Type","application/json")
	_,err = w.Write(core.RefreshToken(requestUser))
	models.CheckError(err,w)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.SignUpDetails)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestUser)
	models.CheckError(err,w)

	//Validating params. Refer models/signupUsers.go
	if _, err := govalidator.ValidateStruct(requestUser); err != nil {
		res := &models.ErrorMessage{err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err,w)
	} else {
		//TODO Add code to save users into DB

		w.Header().Set("Content-Type", "application/json")
		res := &models.SuccessMessage{"Account created successfully"}
		err := json.NewEncoder(w).Encode(res)
		models.CheckError(err,w)
	}
}

//TODO LogoutHandler should be coded
func LogoutHandler(w http.ResponseWriter, r *http.Request) {

}
