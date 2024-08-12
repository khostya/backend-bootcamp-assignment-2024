package cache

import (
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/ds/lfu"
	"sync"
	"time"
)

var _ Cache[uint, domain.House] = (*HouseCache)(nil)

type (
	HouseCache struct {
		cache Cache[uint, domain.House]

		lock sync.RWMutex
	}
)

func NewHouseCache(cap uint, ttl time.Duration) *HouseCache {
	return &HouseCache{
		cache: lfu.NewLFU[uint, domain.House](cap, ttl),
	}
}

func (h *HouseCache) Put(k uint, v domain.House) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.cache.Put(k, v)
}

func (h *HouseCache) Get(k uint) (domain.House, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.cache.Get(k)
}

func (h *HouseCache) Remove(k uint) bool {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.cache.Remove(k)
}
