package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Outbox struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	AggregateType string
	AggregateID   string
	EventType     string
	Payload       datatypes.JSON
	PublishedAt   *time.Time
	CreatedAt     time.Time
}
