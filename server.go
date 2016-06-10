package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sqdron/sqdron-go-auth/middleware"
	"log"
	"os"
)
//
//func api() {
//	return func(c *gin.Context) {
//		c.JSON(200, gin.H{"status": "you are logged in"})
//	}
//}

func main() {
	log.Println("Server is preparing to start")

	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "3000"
	}

	authorized := r.Group("/")

	authorized.Use(middleware.TokenAuthMiddleware())
	{

		authorized.POST("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "you are logged in"})
		})

		authorized.GET("/token", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		authorized.GET("/read", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		//authorized.GET("/test", api);

	}
	r.Run(":" + port);
	//s.Setup()();
	////fb.FacebookConfig{}
	//Application := app.GetApplication()
	//
	////app.UseFacebookAuth();
	//var s = utils.JsonFileSource{};
	//s.Load();
	//log.Println("Server is running")
	//Application.Run()
	//log.Println("Exit")
}