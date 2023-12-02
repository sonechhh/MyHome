package items

type Item struct {
	Type string
}

func (item Item) getItemString() string {
	return "\t\t\t\t" + item.Type + "\n"
}
