package main

import (
	AuthController "panguintz/jwt-api/controller/auth"
	"panguintz/jwt-api/orm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func main() {
	orm.DBConnect()
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}