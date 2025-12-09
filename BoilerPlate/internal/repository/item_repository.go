package repository

import (
    "errors"
    "sync"

    "github.com/example/gin-jwt-swagger/internal/models"
)

var ErrNotFound = errors.New("not found")

type InMemoryRepo struct {
    mu    sync.RWMutex
    store map[int]models.Item
    next  int
}

func NewInMemory() *InMemoryRepo {
    r := &InMemoryRepo{
        store: map[int]models.Item{},
        next:  1,
    }
    return r
}

func (r *InMemoryRepo) List() []models.Item {
    r.mu.RLock()
    defer r.mu.RUnlock()
    res := make([]models.Item, 0, len(r.store))
    for _, v := range r.store {
        res = append(res, v)
    }
    return res
}

func (r *InMemoryRepo) Get(id int) (models.Item, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    it, ok := r.store[id]
    if !ok {
        return models.Item{}, ErrNotFound
    }
    return it, nil
}

func (r *InMemoryRepo) Create(it models.Item) models.Item {
    r.mu.Lock()
    defer r.mu.Unlock()
    it.ID = r.next
    r.next++
    r.store[it.ID] = it
    return it
}

func (r *InMemoryRepo) Update(id int, it models.Item) (models.Item, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, ok := r.store[id]; !ok {
        return models.Item{}, ErrNotFound
    }
    it.ID = id
    r.store[id] = it
    return it, nil
}

func (r *InMemoryRepo) Delete(id int) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, ok := r.store[id]; !ok {
        return ErrNotFound
    }
    delete(r.store, id)
    return nil
}
