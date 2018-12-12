package core

import (
	"../models"
	"encoding/json"
	"net/http"
)

func Login(requestUser *models.User) (int, []byte) {
	authBackend := InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		token, err := authBackend.Generate(requestUser.ID)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(models.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *models.User) []byte {
	authBackend := InitJWTAuthenticationBackend()
	token, err := authBackend.Generate(requestUser.ID)
	if err != nil {
		panic(err)
	}
	response,err := json.Marshal(models.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

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