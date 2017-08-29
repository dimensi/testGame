package main

func main() {

}

func handleCommand(command string) string {
	parsedCommand := parserCommand(command)

	var answer string
	var err *gameError

	switch cmd := parsedCommand.method; cmd {
	case "lookAround":
		answer = player.lookAround(&player.CurrentRoom)
	case "goToRoom":
		answer, err = player.goToRoom(parsedCommand.argument[0])
	case "dressIt":
		answer, err = player.dressIt(parsedCommand.argument[0])
	case "takeIt":
		answer, err = player.takeIt(parsedCommand.argument[0])
	}

	if err != nil {
		return mapErrors[err.message]
	}

	if answer != "" {
		return answer
	}

	return "неизвестная команда"
}

var player *Player

func initGame() {
	player = PlayerCreate()
}
