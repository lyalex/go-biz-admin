// 注册，前端的User信息收集之后，交给本函数，写入数据库
package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lyalex/go-biz-admin/database"
	"github.com/lyalex/go-biz-admin/models"
	"github.com/lyalex/go-biz-admin/utils"
)

func Register(c *gin.Context) {
	/*
		data:
		 "firstname" : "xxx"
		 "lastname" : "xxx"
		 "email" : "xxx"
		 "password" : "xxx"
		 "password_confirmed" : "xxx"
		 "roleid" : "xxx"
	*/
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Input data is not JSON format"})
		return
	}

	if data["password"] != data["password_confirmed"] {
		fmt.Println("password does not match...")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "password does not match"})
		return
	}

	var role_id int
	role_id, _ = strconv.Atoi(data["role"])

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleId:    uint(role_id),
	}

	user.TranslatePassword(data["password"])

	database.DB.Create(&user)
}

func Login(c *gin.Context) {
	var data map[string]string
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Input data is not JSON format"})
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)
	// select * from user where email = ? limit 1 order by id;

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "incorrect password"})
		return
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(user.Id)))

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie("jwt", token, 3600, "", "", false, true)

	c.JSON(http.StatusOK, user)
}

func User(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		fmt.Println("Cookie not set")
	}
	id, _ := utils.ParseJWt(cookie)

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	c.JSON(http.StatusOK, user)
}

func AllUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Role").Find(&users)

	c.JSON(http.StatusOK, users)
}

func FindAUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&user)

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"},
		)
		return
	}

	user.TranslatePassword("1234")

	database.DB.Create(&user)

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"},
		)
		return
	}

	database.DB.Model(&user).Updates(user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "user delete successfully"})
}
