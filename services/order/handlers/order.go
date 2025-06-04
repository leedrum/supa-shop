package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leedrum/supa-shop/services/order/db"
	"github.com/leedrum/supa-shop/services/order/models"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func Create(c *gin.Context) {
	var input struct {
		CustomerID string  `json:"customer_id" binding:"required"`
		Amount     float64 `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID := uuid.New()
	order := models.Order{
		ID:         orderID,
		CustomerID: input.CustomerID,
		Amount:     input.Amount,
		CreatedAt:  time.Now(),
	}

	eventPayload := map[string]interface{}{
		"order_id":    order.ID.String(),
		"customer_id": order.CustomerID,
		"amount":      order.Amount,
		"created_at":  order.CreatedAt,
	}
	payloadBytes, err := json.Marshal(eventPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal event payload"})
		return
	}

	outboxEvent := models.OutboxEvent{
		ID:            uuid.New(),
		AggregateType: "order",
		AggregateID:   order.ID.String(),
		EventType:     "OrderCreated",
		Payload:       datatypes.JSON(payloadBytes),
		CreatedAt:     time.Now(),
	}

	err = db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		if err := tx.Create(&outboxEvent).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order_id": order.ID.String()})
}
