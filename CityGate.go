package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Starting CityGate")
	for i := 0; i < 0xFFFF; i++ {
		fmt.Print(i)
		fmt.Print(' ')
	}
}
