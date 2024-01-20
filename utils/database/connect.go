package database

import (
	"fmt"

	"github.com/lyalex/go-biz-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:12345678@/go_biz_admin"), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection failed\n", db)
		return
	}
	fmt.Println("database init...\n", db)
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
	DB = db
}
