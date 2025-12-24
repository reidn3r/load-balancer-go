package strategies

import (
	"github.com/reidn3r/load-balancer-golang/backend"
)

type RoundRobinStrategy struct {
	index uint64
}

func (rr *RoundRobinStrategy) GetNextBackend(backends []backend.Backend) backend.Backend {
	backendPoolSize := len(backends)
	nextIdx := (rr.index + 1) % uint64(backendPoolSize)
	rr.index = nextIdx
	return backends[nextIdx]
}
