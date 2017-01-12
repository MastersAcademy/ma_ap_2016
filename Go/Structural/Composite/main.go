package main

import "github.com/MastersAcademy/ma_ap_2016/Go/Structural/Composite/component"
import "fmt"

func main() {
	usr := component.NewDirectory("usr")
	local := component.NewDirectory("local")
	file1 := component.NewFile("file1.txt")
	etc := component.NewDirectory("etc")
	file2 := component.NewFile("file2.txt")
	file3 := component.NewFile("file3.txt")

	usr.AddChild(local)
	local.AddChild(file1)
	local.AddChild(etc)
	etc.AddChild(file2)
	etc.AddChild(file3)

	fmt.Printf(usr.Print(""))
}
