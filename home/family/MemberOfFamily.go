package family

import "fmt"

type MemberOfFamily struct {
	Name    string
	Color   string
	Age     int
	Address string
	Food    string
}

func (member MemberOfFamily) getMemberOfFamilyString() string {
	resString := ""
	resString += "\t\tИмя: " + member.Name + "\n" +
		"\t\tЛюбимый цвет: " + member.Color + "\n" +
		"\t\tВозраст: " + fmt.Sprint(member.Age) + "\n" +
		"\t\tАдрес проживания: " + member.Address + "\n" +
		"\t\tЛюбимое блюдо: " + member.Food + "\n"
	return resString
}
