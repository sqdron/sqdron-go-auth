package services

import (
	"models"

	"time"
	jwt_lib "github.com/dgrijalva/jwt-go"
)

var (
	mysupersecretpassword = "unicornsAreAwesome"
)

func Signin(user *models.User)(string, error){

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
//
//
//func RefreshToken(requestUser *models.User) []byte {
//
//	// Create the token
//	token := jwt.New(jwt.SigningMethodHS256)
//	// Set some claims
//	token.Claims["foo"] = "bar"
//	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
//	// Sign and get the complete encoded token as a string
//	tokenString, err := token.SignedString(mySigningKey)
//
//	token := jwt.New(jwt.GetSigningMethod("AppEngine"))
//
//	authBackend := authentication.InitJWTAuthenticationBackend()
//	token, err := authBackend.GenerateToken(requestUser.UUID)
//	if err != nil {
//		panic(err)
//	}
//	response, err := json.Marshal(parameters.TokenAuthentication{token})
//	if err != nil {
//		panic(err)
//	}
//	return response
//}
//
//func Logout(req *http.Request) error {
//	authBackend := authentication.InitJWTAuthenticationBackend()
//	tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
//		return authBackend.PublicKey, nil
//	})
//	if err != nil {
//		return err
//	}
//	tokenString := req.Header.Get("Authorization")
//	return authBackend.Logout(tokenString, tokenRequest)
//}