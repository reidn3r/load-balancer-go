package lb_algorithms

import (
	"net/http"

	"github.com/reidn3r/load-balancer-golang/backend"
)

type LeastConnStrategy struct {
	BackendMapper map[backend.Backend]uint64
	//mutex
}

func (lc *LeastConnStrategy) GetNextBackend(pool []backend.Backend) backend.Backend {
	minConn := pool[0]
	for _, backend := range pool[1:] {
		conn := lc.BackendMapper[backend]
		if conn < lc.BackendMapper[minConn] {
			minConn = backend
		}
	}
	return minConn
}

func (lc *LeastConnStrategy) Serve(
	serverPool []backend.Backend,
	w http.ResponseWriter,
	r *http.Request,
) {
	//add. mutex
	leastConn := lc.GetNextBackend(serverPool)
	lc.BackendMapper[leastConn]++
	leastConn.Proxy.ServeHTTP(w, r)
	lc.BackendMapper[leastConn]--
}
