package entities

type CompanyEntity struct {
	Name          string
	OfficialEmail string
	InfoEmail     string
	MainPhone     string
	FaxPhone      string
}

func (ce *CompanyEntity) GetComplexData() map[string]string {
	m := make(map[string]string)
	m["name"] = ce.Name
	m["officialEmail"] = ce.OfficialEmail
	m["infoEmail"] = ce.InfoEmail
	m["mainPhone"] = ce.MainPhone
	m["faxPhone"] = ce.FaxPhone
	return m
}
