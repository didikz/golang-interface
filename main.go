package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	triangleShape := triangle{
		base:   3,
		height: 5,
	}

	printArea(triangleShape)

	squareShape := square{
		sideLength: 5,
	}

	printArea(squareShape)

	file := os.Args[1]
	txt, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	bs := make([]byte, 9999)
	txt.Read(bs)
	fmt.Println(string(bs))
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}
