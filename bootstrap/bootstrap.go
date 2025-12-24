package bootstrap

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/reidn3r/load-balancer-golang/backend"
	"github.com/reidn3r/load-balancer-golang/config"
	load_balancer "github.com/reidn3r/load-balancer-golang/internal/balancer"
	strategy_factory "github.com/reidn3r/load-balancer-golang/internal/strategies/factory"
)

func Bootstrap(configPath string) {
	cfg, err := readFile(configPath)
	if err != nil {
		msg := fmt.Sprintf("Error while reading config file: %s\n", configPath)
		log.Fatal(msg)
	}
	setup(cfg)
}

func readFile(path string) (config.ConfigObject, error) {
	buffer, err := os.ReadFile(path)
	if err != nil {
		ErrMsg := fmt.Sprintf("Erro ao abrir arquivo: %s\n", path)
		log.Fatal(ErrMsg)
		log.Fatal("Error: ", err)
	}

	var cfgObject config.ConfigObject
	err = json.Unmarshal(buffer, &cfgObject)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", cfgObject)

	return cfgObject, nil //retorna nil
}

func setup(config config.ConfigObject) {
	strategy, err := strategy_factory.BuildStrategy(config)
	if err != nil {
		log.Fatal(err)
	}

	lb := load_balancer.NewLoadBalancer(strategy)

	for _, b := range config.BackendPool {
		lb.AddBackend(*backend.CreateBackend(b.URL))
	}

	addr := fmt.Sprintf(":%d", config.ApplicationPort)
	server := &http.Server{
		Addr:    addr,
		Handler: lb,
	}

	msg := fmt.Sprintf("[LB]: live at http://localhost:%d", config.ApplicationPort)
	fmt.Println(msg)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("Error while creating load balance http server")
	}

	log.Fatal(server.ListenAndServe())
}
