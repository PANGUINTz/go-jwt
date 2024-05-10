package auth

import (
	"fmt"
	"net/http"
	"panguintz/jwt-api/orm"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterType struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}


func hashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([] byte(password), 14)
	return string(bytes), err
}

func verifyPassword(hashpassword string ,password string) bool {
	bytes := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	if bytes != nil {
		fmt.Println("Password mismatch:", bytes)
		return false
	}
	return true
}

func Register(c *gin.Context) {
	var json RegisterType
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userExist := orm.User{}
		orm.Db.Where("username = ?", json.Username).First(&userExist)

		fmt.Print(userExist)

		if userExist.ID > 0{
			c.JSON(http.StatusBadRequest, gin.H{"error": "username is already exists"})
			return
		} 

		encryptPassword, err := hashedPassword(json.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user := orm.User{Username: json.Username, Password: encryptPassword, Fullname: json.Fullname, Avatar: json.Avatar}

		result := orm.Db.Create(&user)

		if result.Error == nil {
			c.JSON(200, gin.H{
				"status": true,
				"message": "Create User Success",
			})
		} else {
			c.JSON(404, gin.H{
				"status": false,
				"message": "User Create Failed",
			})
		}
}

type LoginType struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(context *gin.Context){
	var json LoginType
	if err := context.ShouldBindJSON(&json); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hasUser := orm.User{}

	orm.Db.Where("username", json.Username).First(&hasUser)

	if(hasUser.ID == 0) {
		context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "user not found"})
		return 
	}
	
	if verifyPassword(json.Password, hasUser.Password) {
		context.JSON(200,gin.H{"status": "200", "message": "success"})
	}
}