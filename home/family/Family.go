package family

type Family struct {
	MemberOfFamilyArr []MemberOfFamily
}

func (family Family) FamilyInfo() string {
	resString := ""
	if len(family.MemberOfFamilyArr) > 0 {
		resString += "\tЧлены семьи:\n"
	}
	for i, member := range family.MemberOfFamilyArr {
		resString += member.getMemberOfFamilyString()
		if (i + 1) != len(family.MemberOfFamilyArr) {
			resString += "\t\t++++++++++++++++++\n"
		}
	}
	return resString
}
