package main

import (
	"strings"
)

type Coordinate struct {
	xPos int32
	yPos int32
}

type Rover struct {
	xPos    int32
	yPos    int32
	visited []Coordinate
}

func CreateRover(initalPosition Coordinate) Rover {
	return Rover{
		xPos:    initalPosition.xPos,
		yPos:    initalPosition.yPos,
		visited: []Coordinate{initalPosition},
	}
}

func (rover *Rover) ExecuteCommand(commandString string) {
	splitCommands := strings.Split(commandString, "")

	for _, command := range splitCommands {
		rover.ExecuteDirectionCommand(command)
	}
}

func (rover *Rover) ExecuteDirectionCommand(command string) {
	if command == "N" {
		rover.yPos += 1
	}

	if command == "S" {
		rover.yPos -= 1
	}

	if command == "E" {
		rover.xPos += 1
	}

	if command == "W" {
		rover.xPos -= 1
	}

	newCoordinate := Coordinate{
		xPos: rover.xPos,
		yPos: rover.yPos,
	}

	rover.visited = append(rover.visited, newCoordinate)
}
