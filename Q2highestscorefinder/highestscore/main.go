package main

import (
	"fmt"
	"log"

	"example.com/highest/highest"
)

func main() {
	numbers := []int{76,-1,0,45}
	highestValue, err := highest.FindHighest(numbers)
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println("Highest value:", highestValue)
}
