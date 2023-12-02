package Rooms

import (
	"fmt"
)

type Rooms struct {
	Rooms []Room
}

func (rooms Rooms) countWindows() int {
	windowsSum := 0
	for _, room := range rooms.Rooms {
		windowsSum += room.WindowsCount
	}
	return windowsSum
}

func (rooms Rooms) countSquare() float32 {
	var resSquare float32 = 0
	for _, room := range rooms.Rooms {
		resSquare += room.countSquare()
	}
	return resSquare
}

func (rooms Rooms) countVolume() float32 {
	var resVolume float32 = 0
	for _, room := range rooms.Rooms {
		resVolume += room.countVolume()
	}
	return resVolume
}

func (rooms Rooms) RoomsInfo() string {
	resString := ""
	if len(rooms.Rooms) > 0 {
		resString += "\tКомнаты:\n"
	}
	for i, room := range rooms.Rooms {
		resString += room.getRoomString() + room.Furnitures.FurnitureInfo() + room.Devices.DevicesInfo()
		if (i + 1) != len(rooms.Rooms) {
			resString += "\t\t------------------------\n"
		}
	}
	return resString
}

func (home Rooms) GeneralRoomParameters() string {
	resString := ""
	resString += "\tПлощадь: " + fmt.Sprint(home.countSquare()) + " м^2\n" +
		"\tОбъем: " + fmt.Sprint(home.countVolume()) + " м^3\n" +
		"\tКоличество окон: " + fmt.Sprint(home.countWindows())
	return resString
}
