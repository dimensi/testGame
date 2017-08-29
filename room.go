package main

type Room struct {
	hello     string
	env       string
	nextRooms []string
	items     []string
}

var rooms = map[string]Room{
	"кухня": {
		env:       "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор",
		nextRooms: []string{"коридор"},
	},
	"коридор": {
		hello:     "ничего интересного. можно пройти - кухня, комната, улица",
		nextRooms: []string{"кухня", "комната", "улица"},
	},
	"комната": {
		hello:     "ты в своей комнате. можно пройти - коридор",
		env:       "на столе: ключи, конспекты, на стуле - рюкзак. можно пройти - коридор",
		nextRooms: []string{"коридор"},
		items:     []string{"ключи", "конспекты", "рюкзак"},
	},
}
