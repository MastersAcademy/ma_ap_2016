// Copyright 2016 MastersAcademy.  All rights reserved.
// Use of this source code is governed by a GPLv3
// license that can be found in the LICENSE file.
//  Creational patterns/Builder implementation

package main

import (
	"fmt"
	"./directors"
	"./builders"
)


func main() {
	fmt.Println("Builder pattern Example")

	uDirector := directors.NewComplexDirector(&builders.UserJSONBuilder{}, 54)
	fmt.Println(uDirector.Construct("User Json", "user_json@test.com", "11122211").GetFormatted())

	cDirector := directors.NewComplexDirector(&builders.CompanyKVBuilder{}, 55)
	fmt.Println(cDirector.Construct("Company 1", "company_1@test.com", "company_1_info@test.com", "443344335", "332233225").GetFormatted())

}






