package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	UUID      uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Basket    any
}

type UserProfile struct {
	UUID   uuid.UUID
	Name   string
	Orders []Order
}

type TTLCache[K comparable, T any] struct {
	mu     sync.RWMutex
	ttl    time.Duration
	values map[K]*ttlCacheValue[T]
}

func NewTTLCache[K comparable, T any](ttl time.Duration) *TTLCache[K, T] {
	return &TTLCache[K, T]{
		ttl:    ttl,
		values: make(map[K]*ttlCacheValue[T]),
	}
}

type ttlCacheValue[T any] struct {
	value     T
	createdAt time.Time
}

func newTTLCacheValue[T any](value T) *ttlCacheValue[T] {
	return &ttlCacheValue[T]{
		value:     value,
		createdAt: time.Now(),
	}
}

func (c *TTLCache[K, T]) Put(key K, value T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.values[key] = newTTLCacheValue(value)
}

func (c *TTLCache[K, T]) Get(key K) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	ttlValue, ok := c.values[key]
	if !ok {
		return *new(T), false
	}
	expiredAt := ttlValue.createdAt.Add(c.ttl)

	if time.Now().After(expiredAt) {
		return *new(T), false
	}
	return ttlValue.value, true
}

func main() {
	usersCache := NewTTLCache[uuid.UUID, *UserProfile](time.Minute)

	user := &UserProfile{
		UUID: uuid.New(),
		Name: "name",

		Orders: []Order{
			{
				UUID:   uuid.New(),
				Basket: "basket",

				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	usersCache.Put(user.UUID, user)

	user, ok := usersCache.Get(user.UUID)

	buf, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("user: \n%s", string(buf))
	log.Printf("ok: %t", ok)
}
