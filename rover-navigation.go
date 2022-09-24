package main

import "fmt"

type void struct {}

var member void

func main() {
    roverA, err := CreateRover (
        Coordinate {
            xPos: 0,
            yPos: 2,
        },
    )
    if err != nil {
        panic("Cannot spawn rover out of bounds")
    }

    roverB, err := CreateRover (
        Coordinate {
            xPos: 4,
            yPos: 1,
        },
    )
    if err != nil {
        panic("Cannot spawn rover out of bounds")
    }

    err = roverA.ExecuteCommand("NEESSS")
    if err != nil {
        fmt.Println("Command puts rover out of bounds")
    }
    err = roverB.ExecuteCommand("WWWNNNEEE")
    if err != nil {
        fmt.Println("Command puts rover out of bounds")
    }

    GetOverlap(roverA, roverB)
}

func GetOverlap(roverA *Rover, roverB *Rover){
    set := make(map[Coordinate]*void)

    for _, coordinate := range roverA.visited {
        set[coordinate] = &member
    }

    for _, coordinate := range roverB.visited {
        x := set[coordinate]
        if x != nil {
            fmt.Print("intersection at ")
            fmt.Println(coordinate)
        }
    }
}
