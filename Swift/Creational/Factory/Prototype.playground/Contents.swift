/*:
 Master Academy - Design Patterns implemented in Swift 3.0
 */

/*:
 # Prototype
 The prototype pattern is used to instantiate a new object by copying all of the properties of an existing object, creating an independent clone. This practise is particularly useful when the construction of a new object is inefficient.
 */

import UIKit
import Foundation

struct Platform {
    var name: String
    var version: String
    
    init(name: String, version: String) {
        self.name = name
        self.version = version
        
    }
}

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

let defaultUbisoftGame = Game(developer: "Ubisoft", version: "1.0.0", platforms: [Platform(name: "PC", version: "Win 7"), Platform(name: "PS", version: "3")], engine: "Anvil", price: 59.99)

let assassinsCreed3 = defaultUbisoftGame.clone()
assassinsCreed3.name = "Assasin's Creed 3"
assassinsCreed3.releaseDate = Date()
print("assassinsCreed3 = \(assassinsCreed3.description)")
print()

let assassinsCreed4 = defaultUbisoftGame.clone()
assassinsCreed4.name = "Assasin's Creed 4"
assassinsCreed4.releaseDate = Date()
print("assassinsCreed4 = \(assassinsCreed4.description)")