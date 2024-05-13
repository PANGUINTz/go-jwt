package user

import (
	"net/http"
	"panguintz/jwt-api/orm"
	"strings"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
)

func GET(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	result := strings.Replace(header, "Bearer ", "", -1)

	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(200, gin.H{"status":true, "data": users, "message": "success.", "header": result})
}

func PROFILE(c* gin.Context) {
	userId := c.MustGet("userId")
	var user orm.User
	orm.Db.First(&user, userId)
	c.JSON(http.StatusOK, gin.H{"status": true, "data": user, "message": "success."})
}