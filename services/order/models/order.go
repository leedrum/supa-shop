package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	CustomerID string
	Amount     float64
	CreatedAt  time.Time
}
