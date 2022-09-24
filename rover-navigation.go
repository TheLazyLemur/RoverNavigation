package main

import "fmt"

type void struct {}

var member void

func main() {
    roverA := CreateRover (
        Coordinate {
            xPos: 0,
            yPos: 2,
        },
    )

    roverB := CreateRover (
        Coordinate {
            xPos: 4,
            yPos: 1,
        },
    )

    roverA.ExecuteCommand("NEESSS")
    roverB.ExecuteCommand("WWWNNNEEE")

    GetOverlap(roverA, roverB)
}

func GetOverlap(roverA Rover, roverB Rover){
    set := make(map[Coordinate]*void)

    for _, coordinate := range roverA.visited {
        set[coordinate] = &member
    }

    for _, coordinate := range roverB.visited {
        x := set[coordinate]
        if x != nil {
            fmt.Print("intersection at")
            fmt.Println(coordinate)
        }
    }
}
