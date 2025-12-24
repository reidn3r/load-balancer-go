package lb_algorithms

import (
	"net/http"
	"sync/atomic"

	"github.com/reidn3r/load-balancer-golang/backend"
)

type RoundRobinStrategy struct {
	index atomic.Int64
}

func (rr *RoundRobinStrategy) GetNextBackend(backends []backend.Backend) backend.Backend {
	backendPoolSize := len(backends)
	idx := rr.index.Add(1) % int64(backendPoolSize)
	return backends[idx]
}

func (rr *RoundRobinStrategy) Serve(
	serverPool []backend.Backend,
	w http.ResponseWriter,
	r *http.Request,
) {

	target := rr.GetNextBackend(serverPool)
	target.Proxy.ServeHTTP(w, r)
}
