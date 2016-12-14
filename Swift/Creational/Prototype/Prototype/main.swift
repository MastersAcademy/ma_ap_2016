//
//  main.swift
//  Prototype
//
//  Created by Anton on 13.12.16.
//  Copyright © 2016 None. All rights reserved.
//

import Foundation

let defaultUbisoftGame = Game(developer: "Ubisoft", version: "1.0.0", platforms: [Platform(name: "PC", version: "Win 7"), Platform(name: "PS", version: "3")], engine: "Anvil", price: 59.99)

let assassinsCreed3 = defaultUbisoftGame.clone()
assassinsCreed3.name = "Assasin's Creed 3"
assassinsCreed3.releaseDate = Date()
print(assassinsCreed3)
print()

let assassinsCreed4 = defaultUbisoftGame.clone()
assassinsCreed4.name = "Assasin's Creed 4"
assassinsCreed4.releaseDate = Date()
print(assassinsCreed4)
