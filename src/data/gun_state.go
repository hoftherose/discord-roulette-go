package data

var (
	DefaultLosers      []bool   = []bool{}
	DefaultTurns       []string = []string{}
	DefaultCurrentTurn int      = 0
)

type TableState struct {
	Losers      []bool   `json:"losers"`
	Turns       []string `json:"turns"`
	CurrentTurn int      `json:"current_turn"`
}

var DefaultTableState TableState = TableState{
	DefaultLosers,
	DefaultTurns,
	DefaultCurrentTurn,
}
