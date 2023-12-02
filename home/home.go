package home

import (
	"MyHome/home/Family"
	"MyHome/home/Rooms"
	"MyHome/home/Rooms/Device"
	"MyHome/home/Rooms/Furniture"
	"MyHome/home/Rooms/Furniture/Items"
	"fmt"
)

type Home struct {
	Family        Family.Family
	Rooms         Rooms.Rooms
	FamilySurname string
}

func (home Home) HomeInfo() {
	fmt.Println("\nДом семьи " + home.FamilySurname + ":")
	fmt.Println(home.Rooms.GeneralRoomParameters())
	fmt.Print(home.Family.FamilyInfo())
	fmt.Print(home.Rooms.RoomsInfo())
}

func Make() Home {
	me := Family.MemberOfFamily{
		Name:               "Соня",
		Age:                20,
		Sex:                "Ж",
		RealAddress:        "Волжский бульвар, 1",
		ResidentialAddress: "Монакская, 33",
	}

	mom := Family.MemberOfFamily{
		Name:               "Яна",
		Age:                43,
		Sex:                "Ж",
		RealAddress:        "Волжский бульвар, 1",
		ResidentialAddress: "Миланская, 12",
	}

	dad := Family.MemberOfFamily{
		Name:               "Денис",
		Age:                43,
		Sex:                "М",
		RealAddress:        "Волжский бульвар, 1",
		ResidentialAddress: "Танковая, 2",
	}
	dog := Family.MemberOfFamily{
		Name:               "Ёся",
		Age:                12,
		Sex:                "М",
		RealAddress:        "Волжский бульвар, 1",
		ResidentialAddress: "Калмыкия -",
	}

	MyRoomDevices := Device.Devices{
		[]Device.Device{
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

	bathroomDevices := Device.Devices{
		[]Device.Device{
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

	kitchenDevices := Device.Devices{
		[]Device.Device{
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

	hallwayDevices := Device.Devices{
		[]Device.Device{
			{
				Type:    "Робот-пылесос",
				Name:    "Legee-7",
				Company: "Hobot",
				Length:  33,
				Width:   34,
				Color:   "Blue-Black",
				Count:   1,
			},
		},
	}

	MyRoomFurnituresThingsTable := Items.Items{
		[]Items.Item{
			{
				Type: "Лапма",
			},
			{
				Type: "Коллаж фотографий",
			},
			{
				Type: "Коробка для мелочей",
			},
		},
	}

	MyRoomFurnituresThingsBed := Items.Items{
		[]Items.Item{
			{
				Type: "Комплект плюшевых капибар",
			},
			{
				Type: "Плед",
			},
			{
				Type: "Гирлянда",
			},
		},
	}

	MyRoomFurnituresThingsCupboard := Items.Items{
		[]Items.Item{
			{
				Type: "Книги",
			},
			{
				Type: "Кисти",
			},
			{
				Type: "Хаги-Ваги",
			},
			{
				Type: "Стопарик из Будапешта (невыпитый)",
			},
		},
	}

	MyRoomFurnitures := Furniture.Furnitures{
		[]Furniture.Furniture{
			{
				Type:   "Рабочий стол",
				Length: 150,
				Width:  50,
				Color:  "Brown",
				Items:  MyRoomFurnituresThingsTable,
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
				Items:  MyRoomFurnituresThingsBed,
			},
			{
				Type:   "Стеллаж",
				Length: 170,
				Width:  100,
				Color:  "Black",
				Items:  MyRoomFurnituresThingsCupboard,
			},
		},
	}

	MyRoomFurnituresThingsDressingArea := Items.Items{
		[]Items.Item{
			{
				Type: "Напольная вешалка",
			},
			{
				Type: "Коврик",
			},

			{
				Type: "Зеркало",
			},
		},
	}

	kitchenRoomFurnitures := Furniture.Furnitures{
		[]Furniture.Furniture{
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
				Items:  MyRoomFurnituresThingsDressingArea,
			},
		},
	}

	hallwayFurnituresThingsСloset1 := Items.Items{
		[]Items.Item{
			{
				Type: "Зимняя одежда",
			},
			{
				Type: "Пакет с пакетами для приличных поводов",
			},
		},
	}

	hallwayFurnituresThingsСloset2 := Items.Items{
		[]Items.Item{
			{
				Type: "Обувь",
			},
		},
	}

	hallwayRoomFurnitures := Furniture.Furnitures{
		[]Furniture.Furniture{
			{
				Type:   "Шкаф",
				Length: 200,
				Width:  100,
				Color:  "Brown",
				Items:  hallwayFurnituresThingsСloset1,
			},
			{
				Type:   "Обувница",
				Length: 100,
				Width:  80,
				Color:  "Brown",
				Items:  hallwayFurnituresThingsСloset2,
			},
		},
	}

	MyRoom := Rooms.Room{
		Name:         "Моя комната",
		Width:        3.5,
		Length:       4,
		Height:       2.6,
		WindowsCount: 1,
		Devices:      MyRoomDevices,
		Furnitures:   MyRoomFurnitures,
	}

	kitchen := Rooms.Room{
		Name:         "Кухня",
		Width:        2.5,
		Length:       2.7,
		Height:       2.6,
		WindowsCount: 1,
		Devices:      kitchenDevices,
		Furnitures:   kitchenRoomFurnitures,
	}

	bathroom := Rooms.Room{
		Name:         "Ванная комната",
		Width:        2,
		Length:       2,
		Height:       2.6,
		WindowsCount: 0,
		Devices:      bathroomDevices,
	}

	hallway := Rooms.Room{
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
		Family:        Family.Family{MemberOfFamilyArr: []Family.MemberOfFamily{me, mom, dad, dog}},
		Rooms:         Rooms.Rooms{Rooms: []Rooms.Room{MyRoom, kitchen, bathroom, hallway}},
	}

	return home
}