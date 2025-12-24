package load_balancer

import (
	"net/http"

	backend "github.com/reidn3r/load-balancer-golang/backend"
	"github.com/reidn3r/load-balancer-golang/internal/strategies"
)

type LoadBalancer struct {
	backend           []backend.Backend
	balancingStrategy strategies.LoadBalancerStrategyInterface
}

func NewLoadBalancer(strategy strategies.LoadBalancerStrategyInterface) *LoadBalancer {
	return &LoadBalancer{
		backend:           make([]backend.Backend, 0, 5),
		balancingStrategy: strategy,
	}
}

func (lb *LoadBalancer) AddBackend(backend backend.Backend) {
	lb.backend = append(lb.backend, backend)
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//OBS(reidner): implementação desse método torna
	// a struct LoadBalancer um Handler,
	// capaz de sre usado no server http
	lb.balancingStrategy.Serve(lb.backend, w, r)
}
