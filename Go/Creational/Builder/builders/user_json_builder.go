package builders

import "../entities"
import	"../entities/formatters"
import	"../entities/base"

type UserJSONBuilder struct {
	complexEntity *base.ComplexEntity
}

func (ub *UserJSONBuilder) Init() {
	ub.complexEntity = &base.ComplexEntity{}
	ub.complexEntity.SetFormatter(formatters.NewJSONFormatter("test_"))
}

func (ub *UserJSONBuilder) SetId(id int) {
	ub.complexEntity.Id = id
}

func (ub *UserJSONBuilder) SetDataEntity(data...interface{}) {

	entity := &entities.UserEntity{}

	if 3 > len(data) {
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

		case 1: // email
			param, ok := p.(string)
			if !ok {
				panic("email parameter not type string.")
			}
			entity.Email = param

		case 2: // phone
			param, ok := p.(string)
			if !ok {
				panic("phone parameter not type string.")
			}
			entity.Phone = param
		default:
			panic("Wrong parameter count.")
		}
	}
	ub.complexEntity.DataEntity = entity
}

func (ub *UserJSONBuilder) GetResult() *base.ComplexEntity {
	return ub.complexEntity
}

