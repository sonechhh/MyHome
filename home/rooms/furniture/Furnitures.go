package furniture

type Furnitures struct {
	FurnitureArr []Furniture
}

func (furniture Furnitures) FurnitureInfo() string {
	resString := ""
	if len(furniture.FurnitureArr) > 0 {
		resString += "\t\tМебель:\n"
	}
	for i, furn := range furniture.FurnitureArr {
		resString += furn.getFurnitureString() + furn.Items.ItemsInfo()
		if (i + 1) != len(furniture.FurnitureArr) {
			resString += "\t\t\t---------------\n"
		}
	}
	return resString
}
