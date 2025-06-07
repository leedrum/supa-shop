package consumers

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

// Event struct that matches what order service sends
type OrderCreatedEvent struct {
	OrderID    string  `json:"order_id"`
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	CreatedAt  string  `json:"created_at"`
}

// Handler for order created messages
func handleOrderCreated(event OrderCreatedEvent) {
	log.Printf("Processing order: ID=%s, CustomerID=%s, Amount=%.2f", event.OrderID, event.CustomerID, event.Amount)

	// Example: adjust product stock, trigger analytics, etc.
	// You can add DB logic here if needed.
}

// Consumer group handler
type OrderEventConsumer struct{}

func (c *OrderEventConsumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (c *OrderEventConsumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (c *OrderEventConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Received message: topic=%s key=%s", msg.Topic, string(msg.Key))

		var event OrderCreatedEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("Failed to unmarshal event: %v", err)
			continue
		}

		handleOrderCreated(event)

		// Mark message as processed
		session.MarkMessage(msg, "")
	}
	return nil
}

func Start() {
	brokers := []string{"localhost:9092"}
	groupID := "supashop.product_service"
	topic := "supashop.order_create_events"

	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumerGroup.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Graceful shutdown
	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, os.Interrupt)
		<-sigterm
		cancel()
	}()

	log.Println("Product service is consuming 'order_create_events' topic...")

	for {
		if err := consumerGroup.Consume(ctx, []string{topic}, &OrderEventConsumer{}); err != nil {
			log.Printf("Error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			break
		}
	}

	log.Println("Product service shut down.")
}
