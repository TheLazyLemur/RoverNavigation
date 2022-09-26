package main

import (
	"errors"
	"strings"
)

var roverOutOfBounds = errors.New("Rover went out of bounds")
var commandIsNotValid = errors.New("Command is not valid")

type Coordinate struct {
	xPos int32
	yPos int32
}

type Rover struct {
	xPos    int32
	yPos    int32
	visited []Coordinate
}

func CreateRover(initalPosition Coordinate) (*Rover, error) {

	if initalPosition.xPos > 4 || initalPosition.xPos < 0 || initalPosition.yPos > 4 || initalPosition.yPos < 0 {
		return nil, roverOutOfBounds
	}

	return &Rover{
		xPos:    initalPosition.xPos,
		yPos:    initalPosition.yPos,
		visited: []Coordinate{initalPosition},
	}, nil
}

func (rover *Rover) ExecuteCommand(commandString string) error {
	splitCommands := strings.Split(commandString, "")

	for _, command := range splitCommands {
		err := rover.ExecuteDirectionCommand(command)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rover *Rover) ExecuteDirectionCommand(command string) error {
    err := validateCommand(command)
    if err != nil {
        return err
    }

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

	if rover.xPos > 4 || rover.xPos < 0 || rover.yPos > 4 || rover.yPos < 0 {
		return roverOutOfBounds
	}

	rover.visited = append(rover.visited, newCoordinate)
	return nil
}

func validateCommand(command string) error{
    allowed := []string{"N", "S", "W", "E"}
    isValid := contains(allowed, command)
    if !isValid {
        return commandIsNotValid
    }

    return nil
}

func contains(arr []string, str string) bool {
   for _, a := range arr {
      if a == str {
         return true
      }
   }
   return false
}
