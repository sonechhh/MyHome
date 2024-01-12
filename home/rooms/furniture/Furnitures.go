package furniture

type Furnitures struct {
	FurnitureArr []Furniture
}

func (furniture Furnitures) FurnitureInfo() string {
	resString := ""
	if len(furniture.FurnitureArr) > 0 {
		resString += "\t\tМебель:\n"
	}

	return resString
}
