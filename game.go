package main

type Player struct {
	Package PlayerPackage
}

func PlayerCreate() *Player {
	player := &Player{}

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

type Room struct {
}

type Item interface {
	Do()
}
