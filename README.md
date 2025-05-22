# supa-shop

The structure is below

```
/supa-shop
├── api-gateway/           # The entry point for external API requests
│   └── main.go
│   └── routes/
│   └── middlewares/
│   └── Dockerfile
│
├── services/              # All microservices live here
│   ├── user-service/
│   │   ├── main.go
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── kafka/
│   │   ├── Dockerfile
│   │   └── go.mod
│   ├── order-service/
│   └── ...
│
├── kafka/                 # Kafka docker-compose config, topics setup, schemas, etc.
│   ├── docker-compose.yml
│   └── topics.sh
│
├── proto/                 # If using gRPC or protobuf definitions
│   └── user.proto
│   └── order.proto
│
├── internal/              # Shared internal libraries (e.g., config, logging, kafka utils)
│   ├── config/
│   ├── logger/
│   ├── kafka/
│
├── deployments/           # Kubernetes YAMLs or docker-compose files
│   ├── docker-compose.yml
│   ├── k8s/
│
├── scripts/               # Helper scripts (e.g., bootstrap, migrate)
│
├── README.md
└── go.work                # Use Go workspaces to link all modules together
```
