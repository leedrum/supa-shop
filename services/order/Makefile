## to create a migration file
## migrate create -ext sql -dir db/migrations -seq create_orders_table
POSTGRESQL_URL='postgres://root:123456@localhost:5432/supa-shop-order?sslmode=disable'

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up
.PHONY: migrate-up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down
.PHONY: migrate-down
