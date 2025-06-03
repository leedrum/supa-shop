package main

import (
	"github.com/leedrum/supa-shop/services/auth-service/db"
	"github.com/leedrum/supa-shop/services/auth-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	r.Run(":9001")
}
