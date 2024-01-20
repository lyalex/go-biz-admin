package main

import (
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/lyalex/go-biz-admin/database"
	"github.com/lyalex/go-biz-admin/models"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		var orderItems []models.OrderItem

		for j := -1; j < rand.Intn(4); j++ {
			price := float32(rand.Intn(90) + 10)
			qty := uint(rand.Intn(4) + 1)

			orderItems = append(orderItems, models.OrderItem{
				OrderId:      uint(rand.Intn(5)),
				ProductTitle: faker.Word(),
				Price:        price,
				Quantity:     qty,
			})
		}

		database.DB.Create(&models.Order{
			FirstName:  faker.FirstName(),
			LastName:   faker.LastName(),
			Email:      faker.Email(),
			OrderItems: orderItems,
			UpdatedAt:  time.Unix(faker.RandomUnixTime(), 0).Format("2024-01-17 15:04:05"),
			CreatedAt:  time.Unix(faker.RandomUnixTime(), 0).Format("2024-01-17 15:04:05"),
			Total:      float32(rand.Intn(90) + 10),
		})
	}
}
