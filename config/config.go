package config

import strategies_types "github.com/reidn3r/load-balancer-golang/internal/strategies/types"

type ConfigObject struct {
	ApplicationPort int                            `json:"applicationPort"`
	Strategy        strategies_types.StrategyValue `json:"balancingStrategy"`
	BackendPool     []BackendConfigObject          `json:"backend"`
}

type BackendConfigObject struct {
	URL    string `json:"url"`
	Weight int    `json:"weight"`
}
