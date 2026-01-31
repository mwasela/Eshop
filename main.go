package main

import (
	"Eshop/config"
	"Eshop/controllers"
	"Eshop/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// init DB
	config.ConnectDatabase()

	//public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	//protected routes
	api := r.Group("/api")

	api.Use(middlewares.JWTAuth())
	{
		//
	}

	r.Run(":4000") // listen and serve on 0.0.0.0:4000 (for windows "localhost:4000")
}
