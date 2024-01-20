package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lyalex/go-biz-admin/database"
	"github.com/lyalex/go-biz-admin/models"
)

func AllRoles(c *gin.Context) {
	var roles []models.Role
	database.DB.Find(&roles)

	c.JSON(http.StatusOK, roles)
}

func FindARole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Find(&role)

	c.JSON(http.StatusOK, role)
}

/*
Role Json

	{
	   name: "admin"
	   permissions: [
		"1",
		"2",
		...
		"8"
	   ]
	}
*/
type RoleCreateDTO struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func CreateRole(c *gin.Context) {
	var roleDto RoleCreateDTO
	if err := c.ShouldBindJSON(&roleDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid user JSON file"},
		)
		return
	}

	permissions := make([]models.Permission, len(roleDto.Permissions))

	for idx, permissionId := range roleDto.Permissions {
		id, _ := strconv.Atoi(permissionId)
		permissions[idx] = models.Permission{Id: uint(id)}
	}

	role := models.Role{
		Name:        roleDto.Name,
		Permissions: permissions,
	}

	database.DB.Create(&role)

	c.JSON(http.StatusOK, role)
}

/*
Role JSON

	{
	   "id": "id"
	   name: "admin"
	   permissions: [
		"1",
		"2",
		...
		"8"
	   ]
	}
*/
func UpdateRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	var roleDto RoleCreateDTO

	if err := c.ShouldBindJSON(&roleDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "invalid role JSON file"},
		)
		return
	}

	permissions := make([]models.Permission, len(roleDto.Permissions))

	for idx, permissionId := range roleDto.Permissions {
		id, _ := strconv.Atoi(permissionId)
		permissions[idx] = models.Permission{Id: uint(id)}
	}

	role := models.Role{
		Id:          uint(id),
		Name:        roleDto.Name,
		Permissions: permissions,
	}

	var result struct{}
	database.DB.Table("role_permissions").Where("role_id", id).Delete(&result)

	database.DB.Model(&role).Updates(role)

	c.JSON(http.StatusOK, role)
}

/*
Role JSON

	{
	   "id": "id"
	}
*/
func DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	c.JSON(http.StatusOK, gin.H{"message": "role delete successfully"})
}
