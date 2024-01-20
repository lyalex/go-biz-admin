package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lyalex/go-biz-admin/database"
	"github.com/lyalex/go-biz-admin/models"
)

// request body includes "page":1/2
func AllOrders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	ret := models.Paginate(database.DB, &models.Order{}, page)
	c.JSON(http.StatusOK, ret)
}

type Sales struct {
	Date string `json:"date"`
	Sum  string `json:"sum"`
}

func Chart(c *gin.Context) {
	var sales []Sales
	database.DB.Raw(`
	     SELECT DATE_FORMART(orders.created_at, '%Y-%m-%d') as date, SUM(order_items.price * order_items.quantity) as sum
		 FROM orders
		 JOIN order_items on orders.id = order_items.order_id
		 GROUP BY date
	`).Scan(&sales)
	c.JSON(http.StatusOK, sales)
}
