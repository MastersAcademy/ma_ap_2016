package directors

import "../builders"
import "../entities/base"

type ComplexDirector struct {
	builder builders.ComplexBuilder
}

func NewComplexDirector(builder builders.ComplexBuilder, id int) *ComplexDirector {
	director := &ComplexDirector{}
	director.builder = builder
	director.builder.Init()
	director.builder.SetId(id)

	return director
}

func (cd *ComplexDirector) Construct(data ...interface{}) *base.ComplexEntity {
	cd.builder.SetDataEntity(data...)
	return cd.builder.GetResult();
}
