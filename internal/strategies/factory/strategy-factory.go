package strategy_factory

import (
	"fmt"
	"net/url"

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
		mapper := make(map[*url.URL]uint64)
		return &lb_algorithms.LeastConnStrategy{BackendMapper: mapper}, nil
	}
	if strategy == strategies_types.WeightedRoundRobin {
		pool := lb_algorithms.BuildWrrPool(config.BackendPool)
		return &lb_algorithms.WeightedRoundRobinStrategy{Pool: pool}, nil
	}
	return nil, fmt.Errorf("invalid strategy: %v", strategy)
}
