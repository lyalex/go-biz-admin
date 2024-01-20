package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lyalex/go-biz-admin/database"
	"github.com/lyalex/go-biz-admin/models"
)

func AllPermissions(c *gin.Context) {
	var Permissions []models.Permission
	database.DB.Find(&Permissions)

	c.JSON(http.StatusOK, Permissions)
}

func FindAPermission(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))

	permission := models.Permission{
		Id: uint(id),
	}

	database.DB.Find(&permission)

	c.JSON(http.StatusOK, permission)
}
