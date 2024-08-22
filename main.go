package main

import (
        "fmt"
	"github.com/lyalex/go-biz-admin/database"
	"github.com/lyalex/go-biz-admin/routes"
	// "routes"
)

func main() {
	database.Connect()
	// connect to database
	// bring up route
	r := routes.SetupRouter()
	r.Run(":8080")
        fmt.println("hello world!")
}
