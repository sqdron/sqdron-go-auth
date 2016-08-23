package middleware
//
import (
	//jwt_lib "github.com/dgrijalva/jwt-go"
	//"golang.org/x/oauth2/jwt"
	"github.com/gin-gonic/gin"
	"os"
)

//var (
//	mysupersecretpassword = "unicornsAreAwesome"
//)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")

		if token == "" {
			respondWithError(401, "API token required", c)
			return
		}

		if token != os.Getenv("API_TOKEN") {
			respondWithError(401, "Invalid API token", c)
			return
		}

		c.Next()
	}
}