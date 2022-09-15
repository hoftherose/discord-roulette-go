package data

var (
	DefaultLosers      []bool   = []bool{}
	DefaultTurns       []string = []string{}
	DefaultCurrentTurn int64    = 0

	DefaultChamber        []bool = []bool{}
	DefaultNumChamber     int64  = 6
	DefaultNumBullet      int64  = 1
	DefaultCurrentChamber int64  = 0
)

type TableState struct {
	Losers      []bool   `json:"losers"`
	Turns       []string `json:"turns"`
	CurrentTurn int64    `json:"current_turn"`
}

type GunState struct {
	Chambers       []bool `json:"chambers"`
	NumChamber     int64  `json:"num_chambers"`
	NumBullets     int64  `json:"num_bullets"`
	CurrentChamber int64  `json:"current_chamber"`
}

var DefaultTableState TableState = TableState{
	DefaultLosers,
	DefaultTurns,
	DefaultCurrentTurn,
}

var DefaultGunState GunState = GunState{
	DefaultChamber,
	DefaultNumChamber,
	DefaultNumBullet,
	DefaultCurrentChamber,
}
