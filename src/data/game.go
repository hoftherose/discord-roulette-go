package data

//go:generate mockgen --destination=./../../mocks/game.go github.com/holy-tech/discord-roulette/src/data Game
type Game interface {
	TakeTurn() bool
	IsAccepted() bool
	GameFinished() bool
	GetChannel() string
}

var (
	DefaultGameAccepted bool   = false
	DefaultChannel      string = ""
)

type GameStatus struct {
	Table        Table
	Revolver     Gun
	GameAccepted bool   `json:"game_accepted,omitempty"`
	Channel      string `json:"channel,omitempty"`
}

var DefaultGameStatus GameStatus = GameStatus{
	DefaultGameTable,
	DefaultRevolver,
	DefaultGameAccepted,
	DefaultChannel,
}

func (s *GameStatus) TakeTurn() bool {
	shot := s.Revolver.Shoot()
	currentTurn := s.Table.GetCurrentTurn()
	if shot {
		seating := s.Table.GetSeating()
		s.Table.SetSeating(append(seating[:currentTurn], seating[currentTurn+1:]...))
		s.Table.SetCurrentTurn(currentTurn)
	} else {
		s.Table.SetCurrentTurn(currentTurn + 1)
	}
	return shot
}

func (s *GameStatus) IsAccepted() bool {
	if s.GameAccepted {
		return true
	}
	for _, player := range s.Table.GetSeating() {
		if !player.HasAccepted() {
			return false
		}
	}
	s.GameAccepted = true
	return true
}

func (s *GameStatus) GameFinished() bool {
	return s.Table.NumPlayers() < 2
}

func (s *GameStatus) GetChannel() string {
	return s.Channel
}
