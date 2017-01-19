//
//  main.swift
//  Reflection
//
//  Created by Anton on 18.01.17.
//  Copyright © 2017 None. All rights reserved.
//

import Foundation

class House {
    var address : String = "221B Baker Street"
    var dog : Dog? = Dog()
}

struct Dog {
	let name = "Dog"
    let weight = 5.68
}

class User {
    var name = "Bob"
    var age : Int = 42
    var dogs = [Dog(), Dog()]
	var isMarried = false
    var additionalData : [String : Any] = ["heair" : "brown", "houses" : [House()]]
}

let user = User()

print(Json.jsonString(from: user))



