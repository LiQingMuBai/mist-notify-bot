package global

// BotState 存储每个聊天中的分页状态
type DepositState struct {
	CurrentPage int64
	TotalPages  int64
}
type CostState struct {
	CurrentPage int64
	TotalPages  int64
}

var (
	DepositStates = make(map[int64]*DepositState) // 按ChatID存储状态
	CostStates    = make(map[int64]*CostState)    // 按ChatID存储状态
)
