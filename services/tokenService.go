package services

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"time"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

type TokenService interface {

}

func GenerateJwtToken() string {
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	token.Claims["ID"] = "Christopher"
	token.Claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(mysupersecretpassword))
	if err != nil {
		//c.JSON(500, gin.H{"message": "Could not generate token"})
	}
	return tokenString;
}

