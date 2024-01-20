package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyalex/go-biz-admin/controllers"
	"github.com/lyalex/go-biz-admin/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "Hell World!") })
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)

	// + 权限函数
	r.Use(middleware.IsAuthenticated)

	r.GET("/api/users", controllers.AllUsers)
	r.GET("/api/users/:id", controllers.FindAUser)
	r.POST("/api/users", controllers.CreateUser)
	r.PUT("/api/users/:id", controllers.UpdateUser)
	r.DELETE("/api/users/:id", controllers.DeleteUser)

	r.GET("/api/roles", controllers.AllRoles)
	r.POST("/api/roles", controllers.CreateRole)
	r.GET("/api/roles/:id", controllers.FindARole)
	r.PUT("/api/roles/:id", controllers.UpdateRole)
	r.DELETE("/api/roles/:id", controllers.DeleteRole)

	r.GET("/api/permissions", controllers.AllPermissions)

	return r
}
