package builders

import "../entities/base"

type ComplexBuilder interface {
	Init()
	SetId(in int)
	SetDataEntity(data...interface{})
	GetResult() *base.ComplexEntity
}
