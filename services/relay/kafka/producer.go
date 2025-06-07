package kafka

import (
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/leedrum/supa-shop/services/relay/models"
	"gorm.io/gorm"
)

func NewKafkaSyncProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func ProcessOutboxEvents(db *gorm.DB, producer sarama.SyncProducer, topic string) {
	for {
		var events []models.Outbox

		// Fetch unpublished events
		err := db.
			Where("published_at IS NULL").
			Order("created_at ASC").
			Limit(10).
			Find(&events).Error
		if err != nil {
			log.Printf("Error fetching outbox events: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		for _, event := range events {
			err := publishToKafka(event, producer, topic)
			if err != nil {
				log.Printf("Failed to publish event %s: %v", event.ID, err)
				continue
			}

			// Mark as published
			now := time.Now()
			db.Model(&models.Outbox{}).
				Where("id = ?", event.ID).
				Update("published_at", now)
		}

		time.Sleep(1 * time.Second) // throttle loop
	}
}

func publishToKafka(event models.Outbox, producer sarama.SyncProducer, topic string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(event.AggregateID),
		Value: sarama.ByteEncoder(event.Payload),
		Headers: []sarama.RecordHeader{
			{Key: []byte("event_type"), Value: []byte(event.EventType)},
		},
	}

	_, _, err := producer.SendMessage(msg)
	return err
}
