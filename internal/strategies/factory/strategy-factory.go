package strategy_factory

import (
	"fmt"

	"github.com/reidn3r/load-balancer-golang/config"
	"github.com/reidn3r/load-balancer-golang/internal/strategies"
	lb_algorithms "github.com/reidn3r/load-balancer-golang/internal/strategies/algorithms"
	strategies_types "github.com/reidn3r/load-balancer-golang/internal/strategies/types"
)

func BuildStrategy(config config.ConfigObject) (strategies.LoadBalancerStrategyInterface, error) {
	strategy := config.Strategy

	if strategy == strategies_types.RoundRobin {
		return &lb_algorithms.RoundRobinStrategy{}, nil
	}
	if strategy == strategies_types.LeastConnections {
		return &lb_algorithms.LeastConnStrategy{}, nil
	}
	if strategy == strategies_types.WeightedRoundRobin {
		pool := lb_algorithms.BuildWrrPool(config.BackendPool)
		return &lb_algorithms.WeightedRoundRobinStrategy{Pool: pool}, nil
	}
	return nil, fmt.Errorf("Invalid strategy: %v\n", strategy)
}
