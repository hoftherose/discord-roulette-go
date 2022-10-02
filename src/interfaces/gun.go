package interfaces

//go:generate mockgen --destination=./mocks/gun.go interfaces Gun
type Gun interface {
	ReloadGun(sizeChamber, numBullets int)
	SpinChamber()
	ShuffleChamber()
	Shoot() bool
	NumBulletsLeft() int
	ChamberSize() int
	Seed() int64
	SetSeed(seed int64)
	CurrentChamber() int
	SetCurrentChamber(currChamber int)
	Chamber() []bool
	SetChamber(chamber []bool)
}
