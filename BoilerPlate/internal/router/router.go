package router

import (
    "github.com/example/gin-jwt-swagger/internal/auth"
    "github.com/example/gin-jwt-swagger/internal/config"
    "github.com/example/gin-jwt-swagger/internal/handlers"
    "github.com/example/gin-jwt-swagger/internal/repository"
    "github.com/example/gin-jwt-swagger/internal/service"
    "github.com/gin-gonic/gin"

    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func Setup(r *gin.Engine) {
    // Simple health
    r.GET("/healthz", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // auth - naive login for demo (issue token for any username/password)
    r.POST("/login", func(c *gin.Context) {
        var b map[string]string
        if err := c.BindJSON(&b); err != nil {
            c.JSON(400, gin.H{"error": "invalid body"})
            return
        }
        username := b["username"]
        password := b["password"]
        if username == "" || password == "" {
            c.JSON(400, gin.H{"error": "username & password required"})
            return
        }
        token, err := auth.NewToken([]byte(config.JWTSecret()), username, 24*60*60*1000000000) // 24h in ns
        if err != nil {
            c.JSON(500, gin.H{"error": "could not create token"})
            return
        }
        c.JSON(200, gin.H{"token": token})
    })

    // swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // app wiring
    repo := repository.NewInMemory()
    svc := service.NewItemService(repo)
    h := handlers.NewItemHandler(svc)

    api := r.Group("/api")
    api.Use(auth.JWTMiddleware([]byte(config.JWTSecret())))
    {
        api.GET("/items", h.List)
        api.GET("/items/:id", h.Get)
        api.POST("/items", h.Create)
        api.PUT("/items/:id", h.Update)
        api.DELETE("/items/:id", h.Delete)
    }
}
