package main

import (
	"fmt"

	"github.com/MastersAcademy/ma_ap_2016/Go/Creational/Prototype/Example2/product"
)

func main() {
	x := product.NewProduct(5, "some description")
	fmt.Printf("x: %+v\n", x)
	y := &product.Product{}
	err := product.Clone(x, y) // CAUTION: It can't handle objects with recursive data structs
	if err != nil {
		fmt.Printf("cloning failed: %s\n", err)
		return
	}
	y.SetDescription("other text")
	fmt.Printf("y: %+v\n", y)
	fmt.Printf("x: %+v\n", x)
}
