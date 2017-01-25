package builders

import "../entities/base"
import "../entities/formatters"
import "../entities"

type CompanyKVBuilder struct {
	complexEntity *base.ComplexEntity
}

func (cb *CompanyKVBuilder) Init() {
	cb.complexEntity = &base.ComplexEntity{}
	cb.complexEntity.SetFormatter(formatters.NewKVFormatter("", ""))
}

func (cb *CompanyKVBuilder) SetId(id int) {
	cb.complexEntity.Id = id
}

func (cb *CompanyKVBuilder) SetDataEntity(data...interface{}) {

	entity := &entities.CompanyEntity{}

	if 5 > len(data) {
		panic("Not enough parameters.")
	}

	for i, p := range data {
		switch i {
		case 0: // name
			param, ok := p.(string)
			if !ok {
				panic("name parameter not type string.")
			}
			entity.Name = param

		case 1: // officialEmail
			param, ok := p.(string)
			if !ok {
				panic("email parameter not type string.")
			}
			entity.OfficialEmail = param

		case 2: // infoEmail
			param, ok := p.(string)
			if !ok {
				panic("phone parameter not type string.")
			}
			entity.InfoEmail = param
		case 3: // mainPhone
			param, ok := p.(string)
			if !ok {
				panic("phone parameter not type string.")
			}
			entity.MainPhone = param
		case 4: // faxPhone
			param, ok := p.(string)
			if !ok {
				panic("phone parameter not type string.")
			}
			entity.FaxPhone = param
		default:
			panic("Wrong parameter count.")
		}
	}
	cb.complexEntity.DataEntity = entity
}

func (ub *CompanyKVBuilder) GetResult() *base.ComplexEntity {
	return ub.complexEntity
}
