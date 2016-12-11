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

func main() {
	fmt.Println("wwerqwerqw")

	user := &ComplexEntity{id:1,
		dataEntity: &UserEntity{
			name: "User",
			email: "test@test.com",
			phone: "555333222"}}

	fmt.Println(user.GetComplexData())
	user.SetFormatter(&JSONFormatter{keyPrefix: "test_"})
	fmt.Println(user.GetFormatted())
	user.SetFormatter(NewKVFormatter("", ""))
	fmt.Println(user.GetFormatted())


	company := &ComplexEntity{id:1,
		dataEntity: &CompanyEntity{
			name: "User",
			officialEmail: "test@test.com",
			infoEmail: "test_info@test.com",
			mainPhone: "555333222",
			faxPhone: "555333111",
		}}

	fmt.Println(company.GetComplexData())
	company.SetFormatter(&JSONFormatter{keyPrefix: "test_"})
	fmt.Println(company.GetFormatted())
	company.SetFormatter(NewKVFormatter("", ""))
	fmt.Println(company.GetFormatted())
}






