package core

import (
	"../models"
	"encoding/json"
	"net/http"
)

func Login(requestUser *models.User) (int, []byte) {
	authBackend := InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		id := "12345" //get id from DB
		token, err := authBackend.Generate(id)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(models.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}

	response, _ := json.Marshal(models.ErrorMessage{"User not found"})
	return http.StatusUnauthorized, response
}

//TODO RefreshToken to be optimized
func RefreshToken(requestUser *models.User) []byte {
	id := "12345" //get id from DB
	authBackend := InitJWTAuthenticationBackend()
	token, err := authBackend.Generate(id)
	if err != nil {
		panic(err)
	}
	response,err := json.Marshal(models.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

//TODO Logout func should be coded
//func Logout(r *http.Request) error {
//	authBackend := InitJWTAuthenticationBackend()
//	tokenRequest, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error){
//		return authBackend.PublicKey, nil
//	})
//	if err != nil {
//		return err
//	}
//	tokenString := r.Header.Get("Authorization")
//
//}