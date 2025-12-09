---
noteId: "fb30d6d0d13111f0892ec10cebd47e8d"
tags: []

---

# Gin JWT Swagger Example

This project is a minimal example of a Gin application with:
- JWT auth
- CRUD endpoints (GET/POST/PUT/DELETE)
- Swagger (swaggo) UI
- Dockerfile (multi-stage)
- Jenkinsfile (simple CI pipeline)
- Kubernetes manifests (namespace, deployment, service, ingress)

## Run locally
- Install dependencies: `go mod download`
- Run: `go run ./cmd/server`

## Build Docker
- `docker build -t local/gin-jwt-swagger:latest .`

## Generate Swagger docs
- Install `swag`: `go install github.com/swaggo/swag/cmd/swag@latest`
- `swag init -g ./cmd/server/main.go`

## Kubernetes
- `kubectl apply -f k8s/namespace.yaml`
- Create jwt secret or use provided yaml
- `kubectl apply -f k8s/deployment.yaml -n gin-jwt`
- `kubectl apply -f k8s/service.yaml -n gin-jwt`
- Configure ingress and DNS as needed
