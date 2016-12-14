//
//  Game.swift
//  Prototype
//
//  Created by Anton on 13.12.16.
//  Copyright © 2016 None. All rights reserved.
//

import Foundation

class Game : CustomStringConvertible {
    var name: String?
    var developer: String
    var version: String
    var platforms: [Platform]
    var engine: String
    var price: Double
    var releaseDate: Date?
    
    init(developer: String,
         version: String,
         platforms: [Platform],
         engine: String,
         price: Double) {
        
        self.developer = developer
        self.version = version
        self.platforms = platforms
        self.engine = engine
        self.price = price
    }
    
    func clone() -> Game {
        return Game(developer: self.developer,
                    version: self.version,
                    platforms: self.platforms,
                    engine: self.engine,
                    price: self.price)
    }
    
    var description: String {
        return "Game{name: \(self.name ?? "nil"), developer: \(self.developer), version: \(self.version), platforms: \(self.platforms), engine: \(self.engine), price: \(self.price), releaseDate: \(self.releaseDate?.description ?? "nil")}"

    }
}
