package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccessfullyCreateRover(t *testing.T) {
	coordinate := Coordinate{
		xPos: 0,
		yPos: 3,
	}

	rover, err := CreateRover(coordinate)

	require.Nil(t, err)
	require.NotNil(t, rover)
    require.Equal(t, coordinate.xPos, rover.xPos)
    require.Equal(t, coordinate.yPos, rover.yPos)
    require.Equal(t, coordinate.xPos, rover.visited[0].xPos)
    require.Equal(t, coordinate.yPos, rover.visited[0].yPos)
}


func TestUnsuccessfullyCreateRoverWithNegativeX(t *testing.T) {
	coordinate := Coordinate{
		xPos: -1,
		yPos: 3,
	}

	rover, err := CreateRover(coordinate)

	require.NotNil(t, err)
	require.Nil(t, rover)
}

func TestUnsuccessfullyCreateRoverWithNegativeY(t *testing.T) {
	coordinate := Coordinate{
		xPos: 3,
		yPos: -3,
	}

	rover, err := CreateRover(coordinate)

	require.NotNil(t, err)
	require.Nil(t, rover)
}

func TestSuccessfullyMoveRoverWithCommandString(t *testing.T){
	coordinate := Coordinate{
		xPos: 0,
		yPos: 3,
	}

	rover, _ := CreateRover(coordinate)
    _ = rover.ExecuteCommand("EEESSS")

    require.Equal(t, int32(3), rover.xPos)
    require.Equal(t, int32(0), rover.yPos)
}

func TestUnsuccessfullyMoveRoverWithCommandString(t *testing.T){
	coordinate := Coordinate{
		xPos: 0,
		yPos: 3,
	}

	rover, _ := CreateRover(coordinate)
    err := rover.ExecuteCommand("EEESSSSSS")
    require.NotNil(t, err)
}

func TestSuccessfullyMoveRoverWithCommandGoingNorth(t *testing.T){ 
    coordinate := Coordinate{
		xPos: 0,
		yPos: 3,
	}

	rover, _ := CreateRover(coordinate)
    _ = rover.ExecuteDirectionCommand("N")

    require.Equal(t, int32(0), rover.xPos)
    require.Equal(t, int32(4), rover.yPos)
}

func TestSuccessfullyMoveRoverWithCommandGoingSouth(t *testing.T){ 
    coordinate := Coordinate{
		xPos: 0,
		yPos: 3,
	}

	rover, _ := CreateRover(coordinate)
    _ = rover.ExecuteDirectionCommand("S")

    require.Equal(t, int32(0), rover.xPos)
    require.Equal(t, int32(2), rover.yPos)
}

func TestSuccessfullyMoveRoverWithCommandGoingEast(t *testing.T){ 
    coordinate := Coordinate{
		xPos: 0,
		yPos: 3,
	}

	rover, _ := CreateRover(coordinate)
    _ = rover.ExecuteDirectionCommand("E")

    require.Equal(t, int32(1), rover.xPos)
    require.Equal(t, int32(3), rover.yPos)
}

func TestSuccessfullyMoveRoverWithCommandGoingWest(t *testing.T){ 
    coordinate := Coordinate{
		xPos: 1,
		yPos: 3,
	}

	rover, _ := CreateRover(coordinate)
    _ = rover.ExecuteDirectionCommand("W")

    require.Equal(t, int32(0), rover.xPos)
    require.Equal(t, int32(3), rover.yPos)
}

func TestUnsuccessfullyMoveRoverWithCommandGoingNorth(t *testing.T){ 
    coordinate := Coordinate{
		xPos: 4,
		yPos: 3,
	}

	rover, _ := CreateRover(coordinate)
    err := rover.ExecuteDirectionCommand("N")

    require.Nil(t, err)
}

func TestUnsuccessfullyMoveRoverWithCommandGoingSouth(t *testing.T){ 
    coordinate := Coordinate{
		xPos: 0,
		yPos: 4,
	}

	rover, _ := CreateRover(coordinate)
    err := rover.ExecuteDirectionCommand("S")
    
    require.Nil(t, err)
}

func TestUnsuccessfullyMoveRoverWithCommandGoingEast(t *testing.T){ coordinate := Coordinate{
		xPos: 0,
		yPos: 4,
	}

	rover, _ := CreateRover(coordinate)
    err := rover.ExecuteDirectionCommand("E")

    require.Nil(t, err)
}

func TestunSuccessfullyMoveRoverWithCommandGoingWest(t *testing.T){ coordinate := Coordinate{
		xPos: 0,
		yPos: 3,
	}

	rover, _ := CreateRover(coordinate)
    err := rover.ExecuteDirectionCommand("W")

    require.Nil(t, err)
}
