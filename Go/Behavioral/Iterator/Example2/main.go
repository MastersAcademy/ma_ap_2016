package main

import (
	"fmt"

	"github.com/MastersAcademy/ma_ap_2016/Go/Behavioral/Iterator/Example2/iterator"
)

type vehicle struct {
	maxSpeed  int
	maxPeople int
	maxCargo  int
}

func main() {
	raceCar := &vehicle{
		maxSpeed:  230,
		maxPeople: 1,
		maxCargo:  100,
	}
	truck := &vehicle{
		maxSpeed:  100,
		maxPeople: 3,
		maxCargo:  5000,
	}
	minivan := &vehicle{
		maxSpeed:  150,
		maxPeople: 8,
		maxCargo:  700,
	}
	data := []*vehicle{raceCar, truck, minivan}

	var maxSpeedSum int
	var maxPeopleSum int
	var maxCargoSum int
	for x := range iterator.Iter(data) { // SHOULD BE USED WITH CAUTION
		v := x.(*vehicle)
		maxSpeedSum += v.maxSpeed
		maxPeopleSum += v.maxPeople
		maxCargoSum += v.maxCargo
	}

	fmt.Println("Average max delivering speed:", maxSpeedSum/len(data))
	fmt.Println("People can be delivered:", maxPeopleSum)
	fmt.Println("Cargo can be delivered:", maxCargoSum)
}
