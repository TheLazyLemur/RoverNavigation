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
