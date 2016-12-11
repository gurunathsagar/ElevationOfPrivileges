package corelogic

type Player struct {
	AttackPoints int
	GamePoints   int
	PlayerInfo   PlayerProfile
	playerType   string
}

func createPlayer(playerId string, playerType string) *Player {
	var playerProfile PlayerInfo = getPlayerProfile(playerId)
	player := Player{0, 0, playerProfile, playerType}
	return &player
}
