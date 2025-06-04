package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OutboxEvent struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey"`
	AggregateType string         // e.g., "order"
	AggregateID   string         // order ID
	EventType     string         // e.g., "OrderCreated"
	Payload       datatypes.JSON // JSON payload of the event
	PublishedAt   *time.Time     // null if not published yet
	CreatedAt     time.Time
}
