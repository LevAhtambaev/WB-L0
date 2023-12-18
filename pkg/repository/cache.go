package repository

import (
	"LO/pkg/models"
	"errors"
	"github.com/google/uuid"
	"sync"
)

var (
	ErrNoOrderInCache = errors.New("no order found")
)

type CacheRepositoryImpl struct {
	mu     *sync.RWMutex
	orders map[uuid.UUID]models.OrderJSON
}

func NewCacheRepositoryImpl() *CacheRepositoryImpl {
	return &CacheRepositoryImpl{
		mu:     &sync.RWMutex{},
		orders: make(map[uuid.UUID]models.OrderJSON),
	}
}

func (cr *CacheRepositoryImpl) AddOrder(order models.OrderJSON) error {
	cr.mu.Lock()
	cr.orders[order.OrderUUID] = order
	cr.mu.Unlock()
	return nil
}

func (cr *CacheRepositoryImpl) GetOrderByID(orderUUID uuid.UUID) (models.OrderJSON, error) {
	cr.mu.RLock()
	order, ok := cr.orders[orderUUID]
	cr.mu.RUnlock()
	if !ok {
		return models.OrderJSON{}, ErrNoOrderInCache
	}
	return order, nil
}
