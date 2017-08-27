package main

type Player struct {
}

type Item interface {
	Do()
}

type PlayerPackage struct {
	items map[string]Item
}
