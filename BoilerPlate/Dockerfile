# builder
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache git ca-certificates
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /app ./cmd/server

# final
FROM gcr.io/distroless/static:nonroot
USER nonroot:nonroot
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
