package load_balancer

import (
	backend "github.com/reidn3r/load-balancer-golang/backend"
	"github.com/reidn3r/load-balancer-golang/internal/strategies"
)

type LoadBalancer struct {
	backend           []backend.Backend
	balancingStrategy strategies.LoadBalancerStrategyInterface
	current           uint64
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{backend: make([]backend.Backend, 0, 5), current: 0}
}

func (lb *LoadBalancer) AddBackend(backend backend.Backend) {
	lb.backend = append(lb.backend, backend)
}
