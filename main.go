package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/reidn3r/load-balancer-golang/backend"
	load_balancer "github.com/reidn3r/load-balancer-golang/internal/balancer"
	"github.com/reidn3r/load-balancer-golang/internal/strategies"
)

func main() {
	roundRobin := &strategies.RoundRobinStrategy{}
	lb := load_balancer.NewLoadBalancer(roundRobin)

	lb.AddBackend(*backend.CreateBackend("http://localhost:3000"))
	lb.AddBackend(*backend.CreateBackend("http://localhost:3001"))
	lb.AddBackend(*backend.CreateBackend("http://localhost:3002"))

	server := &http.Server{
		Addr:    ":8080",
		Handler: lb,
	}
	fmt.Println("[LB]: live at http://localhost:8080")
	server.ListenAndServe()
	log.Fatal(server.ListenAndServe())
}
