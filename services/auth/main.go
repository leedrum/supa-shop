package main

import (
	"github.com/leedrum/supa-shop/services/auth/db"
	"github.com/leedrum/supa-shop/services/auth/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)

	r.Run(":9001")
}
