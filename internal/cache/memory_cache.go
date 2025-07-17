package cache

import (
	"context"
	"sync"
	"time"
)

type memoryCache struct {
	store sync.Map
}

type cacheItem struct {
	value      string
	expiration int64
}

func NewMemoryCache() Cache {
	return &memoryCache{}
}

func (m *memoryCache) Get(key string) (string, error) {
	if val, ok := m.store.Load(key); ok {
		item := val.(*cacheItem)
		if item.expiration == 0 || item.expiration > time.Now().UnixNano() {
			return item.value, nil
		}
		m.store.Delete(key)
	}
	return "", nil
}

func (m *memoryCache) Set(key string, value string, expiration time.Duration) error {
	var exp int64
	if expiration > 0 {
		exp = time.Now().Add(expiration).UnixNano()
	}
	m.store.Store(key, &cacheItem{
		value:      value,
		expiration: exp,
	})
	return nil
}

func (m *memoryCache) Delete(key string) error {
	m.store.Delete(key)
	return nil
}

func (m *memoryCache) Exists(key string) (bool, error) {
	_, ok := m.store.Load(key)
	return ok, nil
}

func (m *memoryCache) Clear(ctx context.Context) error {
	m.store = sync.Map{}
	return nil
}

func (m *memoryCache) Close() error {
	m.Clear(context.Background())
	return nil
}
