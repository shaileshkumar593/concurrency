package main

import (
    "log"
    "os"

    "github.com/example/gin-jwt-swagger/internal/router"
    "github.com/gin-gonic/gin"
)

func main() {
    // allow override port via PORT env var
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r := gin.Default()
    router.Setup(r)

    log.Printf("Starting server on :%s\n", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("server error: %v", err)
    }
}
