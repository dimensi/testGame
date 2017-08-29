package main

import "strings"

type Command struct {
	method   string
	argument []string
}

var mapOfMethods = map[string]string{
	"осмотреться": "lookAround",
	"идти":        "goToRoom",
	"одеть":       "dressIt",
	"взять":       "takeIt",
	"применить":   "apply",
}

func parserCommand(command string) *Command {
	sliceStr := strings.Split(command, " ")

	method := mapOfMethods[sliceStr[0]]

	newCommand := &Command{
		method: method,
	}

	if len(sliceStr) > 1 {
		newCommand.argument = sliceStr[1:]
	}

	return newCommand
}

type gameError struct {
	message string
}

var mapErrors = map[string]string{
	"unknownCommandError": "неизвестная команда",
	"noSuchThing":         "нет такого",
	"noSuchRoom":          "нет пути в комната",
	"noWherePut":          "некуда класть",
}

var unknownCommandError = gameError{"unknownCommandError"}
var noSuchThing = gameError{"noSuchThing"}
var noSuchRoom = gameError{"noSuchRoom"}
var noWherePut = gameError{"noWherePut"}
