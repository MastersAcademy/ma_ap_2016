package main

import (
	"github.com/MastersAcademy/ma_ap_2016/Go/Creational/AbstractFactory/example2_abstractFactory/EnchantedMaze"
)

func main(){
	factory := &EnchantedMaze.EnchantedMazeFactory{}
	maze := EnchantedMaze.CreateMaze(factory)
	(*maze).GetRoom(1).Enter()
}


