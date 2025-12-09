package service

import (
    "github.com/example/gin-jwt-swagger/internal/models"
    "github.com/example/gin-jwt-swagger/internal/repository"
)

type ItemService struct {
    repo *repository.InMemoryRepo
}

func NewItemService(r *repository.InMemoryRepo) *ItemService {
    return &ItemService{repo: r}
}

func (s *ItemService) List() []models.Item {
    return s.repo.List()
}

func (s *ItemService) Get(id int) (models.Item, error) {
    return s.repo.Get(id)
}

func (s *ItemService) Create(it models.Item) models.Item {
    return s.repo.Create(it)
}

func (s *ItemService) Update(id int, it models.Item) (models.Item, error) {
    return s.repo.Update(id, it)
}

func (s *ItemService) Delete(id int) error {
    return s.repo.Delete(id)
}
