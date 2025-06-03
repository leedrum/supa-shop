package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/leedrum/supa-shop/services/auth-service/db"
	"github.com/leedrum/supa-shop/services/auth-service/models"
	"github.com/leedrum/supa-shop/services/auth-service/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

const (
	tokenExpireTime = 3600
	bcryptCost      = 10
)

type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var user models.User
	var loginParams LoginParams
	if err := c.ShouldBindJSON(&loginParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	db.DB.Where("username = ?", loginParams.Username).First(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginParams.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.Role, time.Duration(tokenExpireTime))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcryptCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can not generate protected password, Please try another password!"})
	}
	user.Password = string(hashedPassword)

	if err := db.DB.Create(&user).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusConflict, gin.H{"error": "Please try later"})
	}

	// For now, just return 201
	c.JSON(http.StatusCreated, gin.H{"message": "User registered! Let's goooooo to login"})
}
