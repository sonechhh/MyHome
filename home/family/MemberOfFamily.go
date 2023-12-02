package family

import "fmt"

type MemberOfFamily struct {
	Name               string
	Sex                string
	Age                int
	ResidentialAddress string
	RealAddress        string
}

func (member MemberOfFamily) getMemberOfFamilyString() string {
	resString := ""
	resString += "\t\tИмя: " + member.Name + "\n" +
		"\t\tПол: " + member.Sex + "\n" +
		"\t\tВозраст: " + fmt.Sprint(member.Age) + "\n" +
		"\t\tАдрес проживания: " + member.RealAddress + "\n" +
		"\t\tАдрес прописки: " + member.ResidentialAddress + "\n"
	return resString
}
