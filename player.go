package main

type Player struct {
	Package     PlayerPackage
	CurrentRoom Room
}

func PlayerCreate() *Player {
	player := &Player{
		CurrentRoom: Rooms["кухня"],
	}

	return player
}

type PlayerPackage struct {
	items map[string]Item
}

func (p *PlayerPackage) Init() {
	p.items = make(map[string]Item)
}

func (p *PlayerPackage) Add(name string, item Item) {
	p.items[name] = item
}

func (p *PlayerPackage) Remove(name string) {
	delete(p.items, name)
}

func (p *PlayerPackage) Get(name string) Item {
	return p.items[name]
}

func (p *Player) lookAround(room *Room) string {
	return room.env
}

func (p *Player) goToRoom(roomName string) string, UnknownCommandError {
	room, ok := Rooms[roomName]

	if !ok {
		return "", UnknownCommandError
	}
	
	p.CurrentRoom = room

	return room.hello, nil
}

func (p *Player) dressIt(item string) string {
	if (item ===)
}