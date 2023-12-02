package rooms

import (
	"MyHome/home/rooms/device"
	"MyHome/home/rooms/furniture"
	"fmt"
)

type Room struct {
	Name         string
	Width        float32
	Length       float32
	Height       float32
	WindowsCount int
	Devices      device.Devices
	Furnitures   furniture.Furnitures
}

func (room Room) countSquare() float32 {
	return room.Width * room.Length
}

func (room Room) countPerimeter() float32 {
	return (room.Length + room.Width) * 2
}

func (room Room) countVolume() float32 {
	return room.Width * room.Length * room.Height
}

func (room Room) getRoomString() string {
	resString := ""
	resString += "\t\tКомната: " + room.Name + "\n" +
		"\t\tШирина: " + fmt.Sprint(room.Width) + " м \n" +
		"\t\tДлинна: " + fmt.Sprint(room.Length) + " м \n" +
		"\t\tВысота: " + fmt.Sprint(room.Height) + " м \n" +
		"\t\tКоличество окон: " + fmt.Sprint(room.WindowsCount) + "\n"
	return resString
}
