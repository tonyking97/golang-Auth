package core

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"../models"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend()  *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey: getPublicKey(),
		}
	}

	return authBackendInstance
}

//TODO Generating tokeen should be optimized to increase speed
func (backend *JWTAuthenticationBackend) Generate(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * time.Duration(72)).Unix(),
		"iat": time.Now().Unix(),
		"sub": id,
	}
	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		panic(err)
		return "", err
	}

	return tokenString, nil
}

func (backend *JWTAuthenticationBackend) Authenticate(user *models.User) bool {
	//TODO Validation from database should be implemented
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("P@ssw0rd"),10)
	//test

	testUser := models.User{
		Username:"tony",
		Password:string(hashedPassword),
	}
	//test end
	return user.Username == testUser.Username && bcrypt.CompareHashAndPassword([]byte(testUser.Password),[]byte(user.Password)) == nil
}

//TODO Logout func to be coded
//func (backend *JWTAuthenticationBackend) Logout(tokenString string, token *jwt.Token) error {
//
//}

func getPrivateKey() *rsa.PrivateKey {
	dir, _ := os.Getwd()
	privateKeyFile, err := os.Open(dir + "/config/keys/private_key")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	dir, _ := os.Getwd()
	publicKeyFile, err := os.Open(dir + "/config/keys/public_key.pub")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
