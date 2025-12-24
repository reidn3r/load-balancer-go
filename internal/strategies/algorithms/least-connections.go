package lb_algorithms

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/reidn3r/load-balancer-golang/backend"
)

type LeastConnStrategy struct {
	BackendMapper map[*url.URL]uint64
	mutex         *sync.Mutex
}

func (lc *LeastConnStrategy) GetNextBackend(pool []backend.Backend) backend.Backend {
	minConn := pool[0]
	minVal := lc.BackendMapper[minConn.Url]

	for _, backend := range pool[1:] {
		conn := lc.BackendMapper[backend.Url]
		if conn < minVal {
			minConn = backend
			minVal = conn
		}
	}
	return minConn
}

func (lc *LeastConnStrategy) Serve(
	serverPool []backend.Backend,
	w http.ResponseWriter,
	r *http.Request,
) {
	lc.mutex.Lock()
	leastConn := lc.GetNextBackend(serverPool)
	lc.BackendMapper[leastConn.Url]++
	lc.mutex.Unlock()

	leastConn.Proxy.ServeHTTP(w, r)
	lc.BackendMapper[leastConn.Url]--
}
