package data

//go:generate mockgen --destination=./../../mocks/game.go github.com/holy-tech/discord-roulette/src/data Game
type Game interface {
	StartGame()
	TakeTurn()
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

func (s *GameStatus) StartGame() error {
	//TODO implement
	return nil
}

func (s *GameStatus) TakeTurn() (bool, error) {
	//TODO implement
	return false, nil
}

func (s *GameStatus) IsAccepted() bool {
	for _, player := range s.Table.GetSeating() {
		if !player.HasAccepted() {
			return false
		}
	}
	return true
}

func (s *GameStatus) GameFinished() bool {
	return s.Table.NumPlayers() < 2
}

func (s *GameStatus) GetChannel() string {
	return s.Channel
}
