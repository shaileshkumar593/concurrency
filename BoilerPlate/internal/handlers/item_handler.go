package handlers

import (
    "net/http"
    "strconv"

    "github.com/example/gin-jwt-swagger/internal/models"
    "github.com/example/gin-jwt-swagger/internal/service"
    "github.com/gin-gonic/gin"
)

type ItemHandler struct {
    svc *service.ItemService
}

func NewItemHandler(svc *service.ItemService) *ItemHandler {
    return &ItemHandler{svc: svc}
}

func (h *ItemHandler) List(c *gin.Context) {
    c.JSON(http.StatusOK, h.svc.List())
}

func (h *ItemHandler) Get(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    it, err := h.svc.Get(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, it)
}

func (h *ItemHandler) Create(c *gin.Context) {
    var it models.Item
    if err := c.BindJSON(&it); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
        return
    }
    created := h.svc.Create(it)
    c.JSON(http.StatusCreated, created)
}

func (h *ItemHandler) Update(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    var it models.Item
    if err := c.BindJSON(&it); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
        return
    }
    updated, err := h.svc.Update(id, it)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, updated)
}

func (h *ItemHandler) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    if err := h.svc.Delete(id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.Status(http.StatusNoContent)
}
