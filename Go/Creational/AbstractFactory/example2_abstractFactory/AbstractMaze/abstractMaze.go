package AbstractMaze

type MapSite interface {
	Enter()
}

type Room interface {
	MapSite
	GetSide(direction string) MapSite
	SetSide(direction string, side MapSite)
	SetRoomId(roomId int64)
	GetRoomId() int64
}

type Door interface {
	MapSite
	IsOpen() bool
	SetOpen(isOpen bool)
}

type Wall interface {
	MapSite
}

type Maze interface {
	AddRoom(room Room)
	GetRoom(roomId int64) Room
}

type MazeFactory interface {
	CreateMaze() Maze
	CreateRoom(roomId int64) Room
	CreateWall() Wall
	CreateDoor(r1 Room, r2 Room) Door
}