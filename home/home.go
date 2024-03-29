package home

import (
	"MyHome/home/family"
	"MyHome/home/rooms"
	"MyHome/home/rooms/device"
	"MyHome/home/rooms/furniture"

	"fmt"
)

type Home struct {
	Family        family.Family
	Rooms         rooms.Rooms
	FamilySurname string
}

func (home Home) HomeInfo() {
	fmt.Println("\n Квартира " + home.FamilySurname + ":")
	fmt.Println(home.Rooms.GeneralRoomParameters())
	fmt.Print(home.Family.FamilyInfo())
	fmt.Print(home.Rooms.RoomsInfo())
}

func Make() Home {
	me := family.MemberOfFamily{
		Name:    "Соня",
		Age:     20,
		Color:   "Тёмно-синий",
		Address: "Монакская, 33",
		Food:    "Макароны с сыром",
	}

	mom := family.MemberOfFamily{
		Name:    "Яна",
		Age:     43,
		Color:   "Сиреневый",
		Address: "Миланская, 12",
		Food:    "Жареная картошка",
	}

	dad := family.MemberOfFamily{
		Name:    "Денис",
		Age:     43,
		Color:   "Зеленый",
		Address: "Танковая, 2",
		Food:    "Борщ",
	}
	dog := family.MemberOfFamily{
		Name:    "Ёся",
		Age:     12,
		Color:   "Красный",
		Address: "Калмыкия -",
		Food:    "Буквально всё",
	}

	MyRoomDevices := device.Devices{
		[]device.Device{
			{
				Type:    "Ноутбук",
				Name:    "ZenBook",
				Company: "Asus",
				Length:  32,
				Width:   20,
				Color:   "Blue",
				Count:   1,
			},

			{
				Type:    "Принтер",
				Name:    "SmartTank150",
				Company: "HP",
				Length:  39.5,
				Width:   30,
				Color:   "Grey",
				Count:   1,
			},

			{
				Type:    "Увлажнитель воздуха",
				Name:    "увлажнялка",
				Company: "Xiaomi",
				Length:  34,
				Width:   18.5,
				Color:   "White",
				Count:   1,
			},

			{
				Type:    "Кондиционер",
				Name:    "EACS-07HAL/N3",
				Company: "Electrolux",
				Length:  25.5,
				Width:   79.1,
				Color:   "White",
				Count:   1,
			},
		},
	}

	bathroomDevices := device.Devices{
		[]device.Device{
			{
				Type:    "Стиральная машина",
				Name:    "Diamond",
				Company: "Samsung",
				Length:  65,
				Width:   45,
				Color:   "White",
				Count:   1,
			},
			{
				Type:    "Фен",
				Name:    "360",
				Company: "BabyLiss",
				Length:  25,
				Width:   12,
				Color:   "Black",
				Count:   1,
			},
		},
	}

	kitchenDevices := device.Devices{
		[]device.Device{
			{
				Type:    "Телевизор",
				Name:    "H4000",
				Company: "Samsung",
				Length:  43.5,
				Width:   64.3,
				Color:   "Black",
				Count:   1,
			},
			{
				Type:    "Тостер",
				Name:    "TT 330D30",
				Company: "Tefal",
				Length:  18.4,
				Width:   27.5,
				Color:   "White-Yellow",
				Count:   1,
			},
			{
				Type:    "Микроволновка",
				Name:    "MS23K3513A",
				Company: "Samsung",
				Length:  27,
				Width:   49,
				Color:   "Black",
				Count:   1,
			},
		},
	}

	hallwayDevices := device.Devices{
		[]device.Device{
			{
				Type:    "Умная колонка",
				Name:    "Алиса",
				Company: "Яндекс",
				Length:  10,
				Width:   10,
				Color:   "Violet",
				Count:   1,
			},
		},
	}

	MyRoomFurnitures := furniture.Furnitures{
		[]furniture.Furniture{
			{
				Type:   "Рабочий стол",
				Length: 150,
				Width:  50,
				Color:  "Brown",
			},
			{
				Type:   "Комод",
				Length: 90,
				Width:  60,
				Color:  "Blonde wood",
			},
			{
				Type:   "Стул",
				Length: 70,
				Width:  40,
				Color:  "Black",
			},
			{
				Type:   "Кровать",
				Length: 200,
				Width:  150,
				Color:  "Black",
			},
			{
				Type:   "Стеллаж",
				Length: 170,
				Width:  100,
				Color:  "Black",
			},
		},
	}

	kitchenRoomFurnitures := furniture.Furnitures{
		[]furniture.Furniture{
			{
				Type:   "Обеденный стол",
				Length: 100,
				Width:  100,
				Color:  "White",
			},

			{
				Type:   "Стул",
				Length: 65,
				Width:  35,
				Color:  "Brown",
			},
			{
				Type:   "Стеллаж",
				Length: 170,
				Width:  100,
				Color:  "Blonde Wood",
			},
		},
	}

	hallwayRoomFurnitures := furniture.Furnitures{
		[]furniture.Furniture{
			{
				Type:   "Шкаф",
				Length: 200,
				Width:  100,
				Color:  "Brown",
			},
			{
				Type:   "Обувница",
				Length: 100,
				Width:  80,
				Color:  "Brown",
			},
		},
	}

	MyRoom := rooms.Room{
		Name:         "Моя комната",
		Width:        3.5,
		Length:       4,
		Height:       2.6,
		WindowsCount: 1,
		Devices:      MyRoomDevices,
		Furnitures:   MyRoomFurnitures,
	}

	kitchen := rooms.Room{
		Name:         "Кухня",
		Width:        2.5,
		Length:       2.7,
		Height:       2.6,
		WindowsCount: 1,
		Devices:      kitchenDevices,
		Furnitures:   kitchenRoomFurnitures,
	}

	bathroom := rooms.Room{
		Name:         "Ванная комната",
		Width:        2,
		Length:       2,
		Height:       2.6,
		WindowsCount: 0,
		Devices:      bathroomDevices,
	}

	hallway := rooms.Room{
		Name:         "Коридор",
		Width:        3.2,
		Length:       1,
		Height:       2.6,
		WindowsCount: 0,
		Devices:      hallwayDevices,
		Furnitures:   hallwayRoomFurnitures,
	}

	home := Home{
		FamilySurname: "Носовых",
		Family:        family.Family{MemberOfFamilyArr: []family.MemberOfFamily{me, mom, dad, dog}},
		Rooms:         rooms.Rooms{Rooms: []rooms.Room{MyRoom, kitchen, bathroom, hallway}},
	}

	return home
}
