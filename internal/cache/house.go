package cache

import (
	"fmt"
	"github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	"github.com/khostya/backend-bootcamp-assignment-2024/pkg/ds/lfu"
	"sync"
	"time"
)

var _ Cache[HouseKey, domain.House] = (*HouseCache)(nil)

type (
	HouseCache struct {
		cache Cache[string, domain.House]

		lock sync.RWMutex
	}

	HouseKey struct {
		id       uint
		userType domain.UserType
	}
)

func NewHouseKey(ID uint, userType domain.UserType) HouseKey {
	return HouseKey{
		id:       ID,
		userType: userType,
	}
}

func (k HouseKey) String() string {
	return fmt.Sprintf("id=%v user_type=%v", k.id, k.userType)
}

func NewHouseCache(cap uint, ttl time.Duration) *HouseCache {
	return &HouseCache{
		cache: lfu.NewLFU[string, domain.House](cap, ttl),
	}
}

func (h *HouseCache) Put(k HouseKey, v domain.House) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.cache.Put(k.String(), v)
}

func (h *HouseCache) Get(k HouseKey) (domain.House, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.cache.Get(k.String())
}

func (h *HouseCache) Remove(k HouseKey) bool {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.cache.Remove(k.String())
}

func (h *HouseCache) RemoveByID(id uint) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	for _, t := range domain.GetALLUserTypes() {
		h.cache.Remove(HouseKey{
			id:       id,
			userType: t,
		}.String())
	}
}
