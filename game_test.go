package main

import (
	"reflect"
	"testing"
)

func Test_parserCommand(t *testing.T) {

	tests := []struct {
		name string
		args string
		want *Command
	}{
		{
			"case 1",
			"осмотреться",
			&Command{
				method: "lookAround",
			},
		},
		{
			"case 2",
			"идти коридор",
			&Command{
				method:   "goToRoom",
				argument: "коридор",
			},
		},
		{
			"case 3",
			"одеть рюкзак",
			&Command{
				method:   "dressIt",
				argument: "рюкзак",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParserCommand(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parserCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
