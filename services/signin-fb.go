package services

import (
	"models"

	"time"
	jwt_lib "github.com/dgrijalva/jwt-go"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

func SigninFacebook(user *models.User)(string, error){

	// Create the token
	token := jwt_lib.New(jwt_lib.SigningMethodES256)
	// Set some claims
	token.Claims["ID"] = user.Username;
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mysupersecretpassword))
	return tokenString, err
	//if err != nil {
	//	c.JSON(500, gin.H{"message": "Could not generate token"})
	//}
	//c.JSON(200, gin.H{"token": tokenString})
	//return http.StatusUnauthorized, []byte("")
}