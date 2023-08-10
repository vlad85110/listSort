package list

import "sync"

type MyMutex struct {
	mu      sync.Mutex
	locked  bool
}

func (m *MyMutex) Lock() {
	m.mu.Lock()
	m.locked = true
}

func (m *MyMutex) Unlock() {
	m.locked = false
	m.mu.Unlock()
}

func (m *MyMutex) IsLocked() bool {
	return m.locked
}