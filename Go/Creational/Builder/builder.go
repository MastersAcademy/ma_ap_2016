// Copyright 2016 MastersAcademy.  All rights reserved.
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.
//  Creational patterns/Builder implementation

package main

import (
	"fmt"
	"strconv"
)

type BFormatter interface {
	FormatData(complexEntity ComplexEntity) string
}

type JSONFormatter struct {
	keyPrefix string
}

func NewJSONFormatter(prefix string) *JSONFormatter {
	fmt := JSONFormatter{}
	fmt.keyPrefix = prefix
	return &fmt
}

func (jf *JSONFormatter) FormatData(complexEntity ComplexEntity) string {
	rmap := complexEntity.dataEntity.GetComplexData()
	result := "{ id:" + strconv.Itoa(complexEntity.id) + "," + "data:{"
	for k := range rmap {
		result += jf.keyPrefix + k + ":" + rmap[k] + ","
	}
	result = result[0:len(result) - 1]
	result += "}"
	return result
}

type KVFormatter struct {
	sep string
	cor string
}

func NewKVFormatter(sep string, cor string) *KVFormatter {
	fmt := KVFormatter{}
	if sep != "" {
		fmt.sep = sep
	} else {
		fmt.sep = "|"
	}
	if cor != "" {
		fmt.cor = cor
	} else {
		fmt.cor = "="
	}
	return &fmt
}

func (kvf *KVFormatter) FormatData(complexEntity ComplexEntity) string {
	rmap := complexEntity.dataEntity.GetComplexData()
	result := "id" + kvf.cor + strconv.Itoa(complexEntity.id) + kvf.sep
	for k := range rmap {
		result += k + kvf.cor + rmap[k] + kvf.sep
	}
	result = result[0:len(result) - 1]
	return result
}

type DataEntity interface {
	GetComplexData() map[string]string
}

type ComplexEntity struct {
	id         int
	dataEntity DataEntity
	formatter  BFormatter
}

func (ce *ComplexEntity) GetComplexData() map[string]string {
	return ce.dataEntity.GetComplexData()
}

func (ce *ComplexEntity) GetFormatted() string {
	return ce.formatter.FormatData(*ce)
}

func (ce *ComplexEntity) SetFormatter(formatter BFormatter) {
	ce.formatter = formatter
}

type UserEntity struct {
	name  string
	email string
	phone string
}

func (ue *UserEntity) GetComplexData() map[string]string {
	m := make(map[string]string)
	m["name"] = ue.name
	m["email"] = ue.email
	m["phone"] = ue.phone
	return m
}

type CompanyEntity struct {
	name          string
	officialEmail string
	infoEmail     string
	mainPhone     string
	faxPhone      string
	formatter     BFormatter
}

func (ce *CompanyEntity) GetComplexData() map[string]string {
	m := make(map[string]string)
	m["name"] = ce.name
	m["officialEmail"] = ce.officialEmail
	m["infoEmail"] = ce.infoEmail
	m["mainPhone"] = ce.mainPhone
	m["faxPhone"] = ce.faxPhone
	return m
}

type ComplexDirector struct {
	builder ComplexBuilder
}

func NewComplexDirector(builder ComplexBuilder, id int) *ComplexDirector {
	director := &ComplexDirector{}

	director.builder = builder
	director.builder.Init()
	director.builder.SetId(id)

	return director
}

func (cd *ComplexDirector) Construct(data ...interface{}) *ComplexEntity {
	cd.builder.SetDataEntity(data...)
	return cd.builder.GetResult();
}

type ComplexBuilder interface {
	Init()
	SetId(in int)
	SetDataEntity(data...interface{})
	GetResult() *ComplexEntity
}

type UserJSONBuilder struct {
	complexEntity *ComplexEntity
}

func (ub *UserJSONBuilder) Init() {
	ub.complexEntity = &ComplexEntity{}
	ub.complexEntity.SetFormatter(NewJSONFormatter("test_"))
}

func (ub *UserJSONBuilder) SetId(id int) {
	ub.complexEntity.id = id
}

func (ub *UserJSONBuilder) SetDataEntity(data...interface{}) {

	entity := &UserEntity{}

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
			entity.name = param

		case 1: // email
			param, ok := p.(string)
			if !ok {
				panic("email parameter not type string.")
			}
			entity.email = param

		case 2: // phone
			param, ok := p.(string)
			if !ok {
				panic("phone parameter not type int.")
			}
			entity.phone = param
		default:
			panic("Wrong parameter count.")
		}
	}
	ub.complexEntity.dataEntity = entity
}

func (ub *UserJSONBuilder) GetResult() *ComplexEntity {
	return ub.complexEntity
}

type CompanyKVBuilder struct {
	complexEntity *ComplexEntity
}

func (cb *CompanyKVBuilder) Init() {
	cb.complexEntity = &ComplexEntity{}
	cb.complexEntity.SetFormatter(NewKVFormatter("", ""))
}

func (cb *CompanyKVBuilder) SetId(id int) {
	cb.complexEntity.id = id
}

func (cb *CompanyKVBuilder) SetDataEntity(data...interface{}) {

	entity := &CompanyEntity{}

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
			entity.name = param

		case 1: // officialEmail
			param, ok := p.(string)
			if !ok {
				panic("email parameter not type string.")
			}
			entity.officialEmail = param

		case 2: // infoEmail
			param, ok := p.(string)
			if !ok {
				panic("phone parameter not type int.")
			}
			entity.infoEmail = param
		case 3: // mainPhone
			param, ok := p.(string)
			if !ok {
				panic("phone parameter not type int.")
			}
			entity.mainPhone = param
		case 4: // faxPhone
			param, ok := p.(string)
			if !ok {
				panic("phone parameter not type int.")
			}
			entity.faxPhone = param
		default:
			panic("Wrong parameter count.")
		}
	}
	cb.complexEntity.dataEntity = entity
}

func (ub *CompanyKVBuilder) GetResult() *ComplexEntity {
	return ub.complexEntity
}

func main() {
	fmt.Println("Builder pattern Example")

	uDirector := NewComplexDirector(&UserJSONBuilder{}, 54)
	fmt.Println(uDirector.Construct("User Json", "user_json@test.com", "11122211").GetFormatted())

	cDirector := NewComplexDirector(&CompanyKVBuilder{}, 55)
	fmt.Println(cDirector.Construct("Company 1", "company_1@test.com", "company_1_info@test.com", "443344335", "332233225").GetFormatted())

}






