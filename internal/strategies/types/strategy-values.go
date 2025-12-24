package strategies_types

type StrategyValue string

const (
	RoundRobin         StrategyValue = "round-robin"
	LeastConnections   StrategyValue = "least-connections"
	WeightedRoundRobin StrategyValue = "weighted-round-robin"
)
