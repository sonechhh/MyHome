package family

import "fmt"

type MemberOfFamily struct {
	Name    string
	Sex     string
	Age     int
	Address string
	Height  int
}

func (member MemberOfFamily) getMemberOfFamilyString() string {
	resString := ""
	resString += "\t\tИмя: " + member.Name + "\n" +
		"\t\tПол: " + member.Sex + "\n" +
		"\t\tВозраст: " + fmt.Sprint(member.Age) + "\n" +
		"\t\tАдрес проживания: " + member.Address + "\n" +
		"\t\tРост: " + fmt.Sprint(member.Height) + "\n"
	return resString
}
