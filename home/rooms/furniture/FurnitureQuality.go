package furniture

import (
	"fmt"
)

type Furniture struct {
	Type   string
	Length float32
	Width  float32
	Color  string
}

func (furniture Furniture) getFurnitureString() string {
	resString := ""
	resString += "\t\t\tТип мебели: " + furniture.Type + "\n" +
		"\t\t\tШирина: " + fmt.Sprint(furniture.Width) + " м \n" +
		"\t\t\tДлинна: " + fmt.Sprint(furniture.Length) + " м \n" +
		"\t\t\tЦвет: " + furniture.Color + " \n"
	return resString
}
