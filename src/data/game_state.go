package data

var (
	DefaultLosers      []bool   = []bool{}
	DefaultTurns       []string = []string{}
	DefaultCurrentTurn int      = 0

	DefaultChamber        []bool = []bool{}
	DefaultNumChamber     int    = 6
	DefaultNumBullet      int    = 1
	DefaultCurrentChamber int    = 0
)

type TableState struct {
	Losers      []bool   `json:"losers"`
	Turns       []string `json:"turns"`
	CurrentTurn int      `json:"current_turn"`
}

type GunState struct {
	Chambers       []bool `json:"chambers"`
	NumChamber     int    `json:"num_chambers"`
	NumBullets     int    `json:"num_bullets"`
	CurrentChamber int    `json:"current_chamber"`
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
