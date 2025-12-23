package strategies

import "github.com/reidn3r/load-balancer-golang/backend"

type LoadBalancerStrategyInterface interface {
	GetNextBackend(backends []backend.Backend) *backend.Backend
}
