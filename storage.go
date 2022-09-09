package LRUCache

import (
	"github.com/tarruin/LRUCache/interfaces"
	"github.com/tarruin/LRUCache/internal"
	"sync"
)

type Storage struct {
	db interfaces.LRUCache
	mu sync.Mutex
}

func NewStorage(size int) *Storage {
	return &Storage{
		db: internal.NewQueue(size),
	}
}

func (s *Storage) Add(key, value string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.db.Add(key, value)
}

func (s *Storage) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.db.Get(key)
}

func (s *Storage) Remove(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.db.Remove(key)
}
