package interfaces

type GunState interface {
	SetNextChamber()
	SpinChamber(bool)
	ClearChamber()
	getSeed()
}
