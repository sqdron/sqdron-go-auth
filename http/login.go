package http

//import (
//	"net/http"
//	"log"
//)

//
//func Login() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		token := c.Request.FormValue("api_token")
//
//		if token == "" {
//			respondWithError(401, "API token required", c)
//			return
//		}
//
//		if token != os.Getenv("API_TOKEN") {
//			respondWithError(401, "Invalid API token", c)
//			return
//		}
//
//		c.Next()
//	}
//}