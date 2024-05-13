package main

import (
	"fmt"
	AuthController "panguintz/jwt-api/controller/auth"
	UserController "panguintz/jwt-api/controller/user"
	"panguintz/jwt-api/middleware"
	"panguintz/jwt-api/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)
func main() {
	orm.DBConnect()
	r := gin.Default()
	r.Use(cors.Default())

	
    err := godotenv.Load(".env")

    if err != nil {
        fmt.Println("Error loading .env file")
    }

	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)

	authorized := r.Group("/api", middleware.AuthVerify()) 
	authorized.GET("/users", UserController.GET)
	authorized.GET("/profile", UserController.PROFILE)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}