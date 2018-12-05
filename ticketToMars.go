package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("%-18v %-5v %10v %6v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Printf("=====================================\n");

	for count:=10; count>=1; count--{

		speed:= rand.Intn(16)+15
		duration := 62100000/speed/(3600*24)
		price:=36;
		tripType := "One-Way"

		spacelineName:=""
		spaceline:= rand.Intn(3)+1
		switch spaceline{

		case 1: spacelineName=("SpaceX ")
		case 2: spacelineName=("Virgin Galactic ")
		case 3: spacelineName=("Space Adventure ")
		}
		price+=(62100000/16/(3600*24)-duration-7)
		if(rand.Intn(2)+1==2){tripType="Round-trip"; price*=2}

		fmt.Printf("%-18v %-5v %10v %6v\n", spacelineName, duration, tripType, price)

	}
}
