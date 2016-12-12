package main

import (
	"fmt"

	"github.com/MastersAcademy/ma_ap_2016/Go/Creational/Singleton/Example1/singleton"
)

func main() {
	x := singleton.GetInstance()
	x.Number = 5
	fmt.Printf("x: %+v\n", x)

	y := singleton.GetInstance()
	y.Description = "test"
	fmt.Printf("y: %+v\n", y)
	fmt.Printf("x: %+v\n", x)
}
