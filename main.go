package main

import "strings"

func main() {

}

func handleCommand(command string) string {
	parsedCommand := ParserCommand(command)

	var answer string
	var err 
	switch cmd := parsedCommand.method; cmd {
	case "lookAround":
		return player.lookAround(&player.CurrentRoom)
	case "goToRoom":
		return player.goToRoom(parsedCommand.argument[0])
	}

	return "неизвестная команда"
}

var player *Player

func initGame() {
	player = PlayerCreate()
}
