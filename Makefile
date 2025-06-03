api-gateway:
	cd api-gateway && go run main.go

auth:
	cd services/auth && go run main.go

kafka-up:
	docker-compose -f kafka/docker-compose.yml up -d

kafka-down:
	docker-compose -f kafka/docker-compose.yml down -d

.PHONY: api-gateway auth kafka-up kafka-down
