package main

import (
	"github.com/leedrum/supa-shop/services/order/db"
	"github.com/leedrum/supa-shop/services/order/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	r.POST("/create", handlers.Create)

	r.Run(":9002")
}
