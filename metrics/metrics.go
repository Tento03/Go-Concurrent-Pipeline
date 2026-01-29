package metrics

import "sync"

type Metrics struct {
	mu        sync.RWMutex
	Processed int
	Failed    int
}

func (m *Metrics) IncProcessed() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Processed++
}

func (m *Metrics) IncFailed() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Failed++
}

func (m *Metrics) Snapshot() (processed, failed int) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.Processed, m.Failed
}
