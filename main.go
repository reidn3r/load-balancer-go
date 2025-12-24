package main

import (
	"log"
	"os"

	"github.com/reidn3r/load-balancer-golang/bootstrap"
	"github.com/reidn3r/load-balancer-golang/config"
)

func main() {
	argData, err := config.ReadArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	bootstrap.Bootstrap(argData.FilePath)
}
