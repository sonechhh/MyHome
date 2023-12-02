package items

type Items struct {
	ItemArr []Item
}

func (items Items) ItemsInfo() string {
	resString := ""
	if len(items.ItemArr) > 0 {
		resString += "\t\t\tВещи:\n"
	}
	for _, furn := range items.ItemArr {
		resString += furn.getItemString()
	}
	return resString
}
