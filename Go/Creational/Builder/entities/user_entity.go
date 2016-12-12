package entities


type UserEntity struct {
	Name  string
	Email string
	Phone string
}

func (ue *UserEntity) GetComplexData() map[string]string {
	m := make(map[string]string)
	m["name"] = ue.Name
	m["email"] = ue.Email
	m["phone"] = ue.Phone
	return m
}
