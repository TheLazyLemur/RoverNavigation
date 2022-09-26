package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	intX, err := getIntInput("RoverA X spawn point")
	if err != nil {
		fmt.Println(err.Error())
        return
	}

	intY, err := getIntInput("RoverA Y spawn point")
	if err != nil {
		fmt.Println(err.Error())
        return
	}

	roverA, err := CreateRover(
		Coordinate{
			xPos: int32(intX),
			yPos: int32(intY),
		},
	)
	if err != nil {
		fmt.Println(err.Error())
        return
	}

	intX, err = getIntInput("RoverB X spawn point")
	if err != nil {
		fmt.Println(err.Error())
        return
	}

	intY, err = getIntInput("RoverB Y spawn point")
	if err != nil {
		fmt.Println(err.Error())
        return
	}

	roverB, err := CreateRover(
		Coordinate{
			xPos: int32(intX),
			yPos: int32(intY),
		},
	)
	if err != nil {
		fmt.Println(err.Error())
        return
	}

	roverACommandString := getStringInput("Command string of directions for RoverA")
	err = roverA.ExecuteCommand(roverACommandString)
	if err != nil {
		fmt.Println(err.Error())
        return
	}

	roverBCommandString := getStringInput("Command string of directions for RoverB")
	err = roverB.ExecuteCommand(roverBCommandString)
	if err != nil {
		fmt.Println(err.Error())
        return
	}

	GetOverlap(roverA, roverB)
}

func GetOverlap(roverA *Rover, roverB *Rover) {
	set := make(map[Coordinate]bool)

	for _, coordinate := range roverA.visited {
		set[coordinate] = true
	}

	for _, coordinate := range roverB.visited {
		x := set[coordinate]
		if x {
			fmt.Print("intersection at ")
			fmt.Println(coordinate)
		}
	}
}

func getIntInput(inputMessage string) (int32, error) {
	fmt.Println(inputMessage)

	var userInput string
	fmt.Scanln(&userInput)
	value, err := strconv.Atoi(userInput)
	if err != nil {
        return 0, errors.New("input should be of type int")
	}

	return int32(value), nil
}

func getStringInput(inputMessage string) string {
	fmt.Println(inputMessage)

	var userInput string
	fmt.Scanln(&userInput)

	return userInput
}
