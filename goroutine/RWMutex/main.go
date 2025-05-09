package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m map[string]string
}

func (s *SafeMap) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.m[key]
}

func (s *SafeMap) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func main() {
	s := &SafeMap{m: make(map[string]string)}
	var wg sync.WaitGroup

	// 書き込み
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.Set("foo", "bar")
	}()

	// 読み込み
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("foo:", s.Get("foo"))
	}()

	wg.Wait()
}