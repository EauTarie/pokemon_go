package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	min := 1
	max := 4
	var randomNumber, xTimesOne, xTimesTwo, xTimesThree int
	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate random number and print on console
	for i := 0; i < 30; i++ {
		randomNumber = rand.Intn(max-min) + min
		fmt.Println(randomNumber)
		switch randomNumber {
		case 1:
			xTimesOne += 1
		case 2:
			xTimesTwo += 1
		default:
			xTimesThree += 1
		}
	}
	fmt.Println("Zapdos :",xTimesOne,"// Articuno",xTimesTwo,"// Moltres",xTimesThree)
}
