package main

type Player struct {
	Package     playerPackage
	CurrentRoom Room
}

func PlayerCreate() *Player {
	player := &Player{
		CurrentRoom: rooms["кухня"],
	}

	return player
}

type playerPackage struct {
	items   map[string]Item
	dressed bool
}

func (p *playerPackage) init() {
	p.items = make(map[string]Item)
	p.dressed = true
}

func (p *playerPackage) add(name string, item Item) {
	p.items[name] = item
}

func (p *playerPackage) remove(name string) {
	delete(p.items, name)
}

func (p *playerPackage) get(name string) Item {
	return p.items[name]
}

func (p *Player) lookAround(room *Room) string {
	return room.env
}

func (p *Player) goToRoom(roomName string) (string, *gameError) {
	room, ok := rooms[roomName]

	if !ok {
		return "", &noSuchRoom
	}

	p.CurrentRoom = room

	return room.hello, nil
}

func (p *Player) dressIt(item string) (string, *gameError) {
	if item == "рюкзак" {
		p.Package.init()
		return "вы одели: " + item, nil
	}

	return "", &unknownCommandError
}

func (p *Player) takeIt(itemStr string) (string, *gameError) {

	if !p.Package.dressed {
		return "", &noWherePut
	}

	item, ok := items[itemStr]

	if !ok {
		return "", &noSuchThing
	}

	p.Package.add(itemStr, item)

	return "предмет добавлен в инвентарь: " + itemStr, nil
}
