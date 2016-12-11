package main

import (
	"fmt"

	"github.com/MastersAcademy/ma_ap_2016/Go/Creational/Prototype/Example1/product"
)

func main() {
	x := product.NewProduct(5, "some description")
	fmt.Printf("x: %+v\n", x)
	y := x.Clone()
	y.SetDescription("other text")
	fmt.Printf("y: %+v\n", y)
	fmt.Printf("x: %+v\n", x)
}
