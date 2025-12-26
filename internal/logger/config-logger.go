package logger

import (
	"log"

	"github.com/reidn3r/load-balancer-golang/config"
)

func LogConfig(cfg config.ConfigObject) {
	log.Println("=== Load Balancer Configuration ===")
	log.Printf("Strategy: %s\n", cfg.Strategy)

	log.Println("Backends:")
	for i, b := range cfg.BackendPool {
		if b.Weight > 0 {
			log.Printf("  %d) URL: %s | Weight: %d\n", i+1, b.URL, b.Weight)
		} else {
			log.Printf("  %d) URL: %s\n", i+1, b.URL)
		}
	}
	log.Println("===================================")
}
