package lru

import (
	"backend-bootcamp-assignment-2024/pkg/ds/linkedlist"
	"time"
)

type item[K comparable, V any] struct {
	expiredAt time.Time

	node *linkedlist.Node[K, V]
}

func (i item[K, V]) Expired(now time.Time) bool {
	return i.expiredAt.Before(now)
}
