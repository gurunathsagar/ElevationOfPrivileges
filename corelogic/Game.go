package corelogic

type Game struct {
	GameId   string
	Player_1 Player
	Player_2 Player
	Board    GameBoard
}

func getUniqueGameId() string {
	return "123"
}

func createGame() *Game {
	var gameId string = getUniqueGameId()

	var game Game = Game{"Matt", "Aimonetti"}
	return &game
}
