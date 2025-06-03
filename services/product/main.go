package main

import (
	"github.com/leedrum/supa-shop/services/product/db"
	"github.com/leedrum/supa-shop/services/product/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	r.POST("/create", handlers.Create)
	r.DELETE("/delete/:id", handlers.Delete)

	r.Run(":9002")
}
