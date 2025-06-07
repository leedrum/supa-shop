package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leedrum/supa-shop/services/relay/kafka"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "123456"),
		getEnv("DB_NAME", "supa-shop-order"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	brokers := []string{"localhost:9092"}
	producer, err := kafka.NewKafkaSyncProducer(brokers)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	log.Println("Starting outbox relay...")
	kafka.ProcessOutboxEvents(db, producer, "supashop.order_create_events")
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
