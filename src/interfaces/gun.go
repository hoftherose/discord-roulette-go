package interfaces

//go:generate mockgen --destination=./../../mocks/gun.go github.com/holy-tech/discord-roulette/src/interfaces Gun
type Gun interface {
	ReloadGun(sizeChamber, numBullets int)
	SpinChamber()
	ShuffleChamber()
	Shoot() bool
	GetNumBulletsLeft() int
	ChamberSize() int
	GetSeed() int64
	SetSeed(seed int64)
	GetCurrentChamber() int
	SetCurrentChamber(currChamber int)
	GetChamber() []bool
	SetChamber(chamber []bool)
}
