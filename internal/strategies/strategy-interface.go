package strategies

import (
	"net/http"

	"github.com/reidn3r/load-balancer-golang/backend"
)

type LoadBalancerStrategyInterface interface {
	GetNextBackend(backends []backend.Backend) backend.Backend
	Serve(pool []backend.Backend, w http.ResponseWriter, r *http.Request)
}
