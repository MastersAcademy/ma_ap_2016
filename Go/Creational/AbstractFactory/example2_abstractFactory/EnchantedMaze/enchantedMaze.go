package EnchantedMaze

import (
	"github.com/MastersAcademy/ma_ap_2016/Go/Creational/AbstractFactory/example2_abstractFactory/AbstractMaze"
	"fmt"
)

type EnchantedRoom struct {
	Directions map[string]AbstractMaze.MapSite
	ID int64
}

func (r EnchantedRoom) GetSide(direction string) AbstractMaze.MapSite{
	return r.Directions[direction]
}

func (r EnchantedRoom) SetSide(direction string, side AbstractMaze.MapSite){
	r.Directions[direction] = side
}

func (r EnchantedRoom) SetRoomId(roomId int64){
	r.ID = roomId
}

func (r EnchantedRoom) GetRoomId() int64{
	return r.ID
}

func (r EnchantedRoom) Enter(){
	fmt.Printf("You enter room %v\n", r.ID)
}

func NewEnchantedRoom(id int64) AbstractMaze.Room{
	return EnchantedRoom{
		ID: id,
		Directions: make(map[string]AbstractMaze.MapSite),
	}
}

type EnchantedDoor struct {
	state bool
	Room1 AbstractMaze.Room
	Room2 AbstractMaze.Room
}

func (d EnchantedDoor)SetOpen(isOpen bool)  {
	d.state = isOpen
}

func (d EnchantedDoor)IsOpen() bool {
	return d.state
}

func (d EnchantedDoor)Enter()  {
	if d.state {
		println("You go throw the door")
	} else {
		println("You can't go throw closed door")
	}
}

func NewEnchantedDoor(r1 AbstractMaze.Room, r2 AbstractMaze.Room) AbstractMaze.Door{
	return EnchantedDoor{
		state: false,
		Room1: r1,
		Room2: r2,
	}
}

type EnchantedWall struct {
}

func (w EnchantedWall)Enter(){
	println("You can't go throw the wall")
}

func NewEnchantedWall() AbstractMaze.Wall {
	return EnchantedWall{}
}

type EnchantedMaze struct {
	rooms map[int64]AbstractMaze.Room
}

func (m *EnchantedMaze) AddRoom(room AbstractMaze.Room){
	m.rooms[room.GetRoomId()] = room
}

func (m *EnchantedMaze) GetRoom(roomId int64) AbstractMaze.Room {
	room, ok := m.rooms[roomId]
	if ok {
		return room
	} else {
		return nil
	}
}

func NewEnchantedMaze() AbstractMaze.Maze{
	return &EnchantedMaze{
		rooms: make(map[int64]AbstractMaze.Room),
	}
}

type EnchantedMazeFactory struct {
}

func (f EnchantedMazeFactory) CreateMaze() AbstractMaze.Maze {
	return NewEnchantedMaze()
}

func (f EnchantedMazeFactory) CreateRoom(roomId int64) AbstractMaze.Room{
	return NewEnchantedRoom(roomId)
}

func (f EnchantedMazeFactory) CreateWall() AbstractMaze.Wall{
	return NewEnchantedWall()
}

func (f EnchantedMazeFactory) CreateDoor(r1 AbstractMaze.Room, r2 AbstractMaze.Room) AbstractMaze.Door{
	return NewEnchantedDoor(r1, r2)
}

func CreateMaze(factory AbstractMaze.MazeFactory) *AbstractMaze.Maze {
	aMaze := factory.CreateMaze()
	room1 := factory.CreateRoom(1)
	room2 := factory.CreateRoom(2)
	aDoor := factory.CreateDoor(room1, room2)
	aMaze.AddRoom(room1)
	aMaze.AddRoom(room2)
	room1.SetSide("North", factory.CreateWall())
	room1.SetSide("East", aDoor)
	room1.SetSide("South", factory.CreateWall())
	room1.SetSide("West", factory.CreateWall())
	room2.SetSide("North", factory.CreateWall())
	room2.SetSide("East", factory.CreateWall())
	room2.SetSide("South", factory.CreateWall())
	room2.SetSide("West", aDoor)
	return &aMaze
}