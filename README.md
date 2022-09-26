# Rover Navigation

## How to build

- Install Golang 1.19
- Once installed run `go mod tidy`
- To run the project `go run .`
- To run the tests `go test -v -cover ./...`

## Example

Consider the given example where Rover A spawns at `X: 0, Y: 2` and Rover B spawns at `X: 4, Y:1`

Rover A input -> NEESSS
Rover B input -> WWWNNNEEE

The below image is a visual representation of the above example

![Example](/images/rover-example.svg)
