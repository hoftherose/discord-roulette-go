package data

var (
	DefaultLosers []bool   = []bool{}
	DefaultTurns  []string = []string{}

	DefaultChamber    []bool = []bool{}
	DefaultNumChamber int64  = 6
	DefaultNumBullet  int64  = 1
)

type TableState struct {
	Losers []bool   `json:"losers"`
	Turns  []string `json:"current_turn"`
}

type GunState struct {
	Chambers   []bool `json:"chambers"`
	NumChamber int64  `json:"num_chambers"`
	NumBullets int64  `json:"num_bullets"`
}

var DefaultTableState TableState = TableState{
	DefaultLosers,
	DefaultTurns,
}

var DefaultGunState GunState = GunState{
	DefaultChamber,
	DefaultNumChamber,
	DefaultNumBullet,
}
