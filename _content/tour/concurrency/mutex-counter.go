//go:build OMIT

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter безопасен для параллельного использования.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc увеличивает счетчик для данного ключа.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Блокировка, чтобы только одна горутина в момент времени могла обращаться к хэш-таблице c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value возвращает текущее значение счетчика для данного ключа.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Блокировка, чтобы только одна горутина в момент времени могла обращаться к хэш-таблице c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
