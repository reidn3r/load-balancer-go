package lb_algorithms

import (
	"net/http"

	"github.com/reidn3r/load-balancer-golang/backend"
	"github.com/reidn3r/load-balancer-golang/config"
)

type WeightedRoundRobinStrategy struct {
	Pool  []*WrrServer
	index uint
}

type WrrServer struct {
	Server backend.Backend
	Weight uint
	count  uint
}

func (wrr *WeightedRoundRobinStrategy) GetNextBackend(pool []backend.Backend) backend.Backend {
	current := wrr.Pool[wrr.index]
	current.count++

	if current.count >= current.Weight {
		current.count = 0
		wrr.index = (wrr.index + 1) % uint(len(wrr.Pool))
	}

	return current.Server
}

func (wrr *WeightedRoundRobinStrategy) Serve(
	serverPool []backend.Backend,
	w http.ResponseWriter,
	r *http.Request,
) {
	server := wrr.GetNextBackend(serverPool)
	server.Proxy.ServeHTTP(w, r)
}

func CreateNewWrrServer(server backend.Backend, weight uint) *WrrServer {
	if weight == 0 {
		weight = 1
	}

	return &WrrServer{
		Server: server,
		Weight: weight,
		count:  0,
	}
}

func BuildWrrPool(backendPool []config.BackendConfigObject) []*WrrServer {

	pool := make([]*WrrServer, 0, len(backendPool))

	for _, backendObj := range backendPool {
		b := backend.CreateBackend(backendObj.URL)
		wrr := CreateNewWrrServer(*b, uint(backendObj.Weight))
		pool = append(pool, wrr)
	}

	return pool
}
