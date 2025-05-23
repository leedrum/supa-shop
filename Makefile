api-gateway:
	cd api-gateway && go run main.go

auth-service:
	cd services/auth-service && go run main.go

kafka-up:
	docker-compose -f kafka/docker-compose.yml up -d

kafka-down:
	docker-compose -f kafka/docker-compose.yml down -d

.PHONY: api-gateway auth-service kafka-up kafka-down
