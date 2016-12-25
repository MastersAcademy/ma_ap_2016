package main

import (
	"fmt"

	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Flyweight/runeFlyweight"
)

func main() {
	factory := runeFlyweight.NewFactory()
	s := "Hello Golang World!"
	for _, r := range []rune(s) {
		fmt.Printf("%v", factory.GetFlyweight(r))
	}
	fmt.Println()
	fmt.Println("Unique runes number:", factory.UniqueRunesNumber())
}
