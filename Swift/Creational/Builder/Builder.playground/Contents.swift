/*:
 Master Academy - Design Patterns implemented in Swift 3.0
 */
/*:
 # Builder pattern

The builder pattern is used to create complex objects with constituent parts that must be created in the same order or using a specific algorithm. An external class controls the construction algorithm.
 
 */

import UIKit
import Foundation

//: 1. Simple FragranceBuilder
class Fragrance: CustomStringConvertible {
    let name: String
    let ingFlower: String
    let ingFruit: String
    let ingWood: String
    let ingSpecial: String
    let powerLevel: Int
    var description: String {
        return "Fragrance:\(name)\n" +
            "Flower:\(ingFlower)\n" +
            "Fruit: \(ingFruit)\n" +
            "Wood: \(ingWood)\n" +
            "Special Ingredience: \(ingSpecial)\n" +
        "Power Level: \(powerLevel)\n"
    }
    
    init(name:String, ingFlower: String, ingFruit: String, ingWood: String, ingSpecial:String, powerLevel: Int) {
        self.name = name
        self.ingFlower = ingFlower
        self.ingFruit = ingFruit
        self.ingWood = ingWood
        self.ingSpecial = ingSpecial
        self.powerLevel = powerLevel
    }
}

class FragranceBuilder {
    fileprivate var ingFlower = "Jasmine"
    fileprivate var ingFruit = "Strawberry"
    fileprivate var ingWood = "Sandalwood"
    fileprivate var ingSpecial = "Fish's tear"
    fileprivate var powerLevel = 0
    func setIngFlower(ingFlower: String){
        self.ingFlower = ingFlower
    }
    func setIngFruit(ingFruit: String){
        self.ingFruit = ingFruit
    }
    func setIngWood(ingWood: String){
        self.ingWood = ingWood
    }
    func setIngSpecial(ingSpecial: String){
        self.ingSpecial = ingSpecial
    }
    func setPowerLevel(powerLevel: Int){
        self.powerLevel = powerLevel
    }
    
    func buildFragrance(name: String) -> Fragrance {
        return Fragrance(name: name, ingFlower: ingFlower, ingFruit: ingFruit, ingWood: ingWood, ingSpecial: ingSpecial, powerLevel: powerLevel)
    }
}

let builder = FragranceBuilder()
builder.setIngFlower(ingFlower: "Rose")
builder.setIngSpecial(ingSpecial: "First Raindrop in Spring")
let roseFragrance = builder.buildFragrance(name: "Rose in Spring")
print(roseFragrance.description)


//: 2. Use specialized FragranceBuilderFactory to build different fragrance
enum FragranceType {
    case Floral
    case Fruity
    case Oceanic
}

class FragranceBuilderFactory {
    fileprivate var ingFlower = "Jasmine"
    fileprivate var ingFruit = "Strawberry"
    fileprivate var ingWood = "Sandalwood"
    fileprivate var ingSpecial = "Fish's tear"
    fileprivate var powerLevel = 0
    fileprivate init(){
        // do nothing
    }
    func setIngFlower(ingFlower: String){
        self.ingFlower = ingFlower
    }
    func setIngFruit(ingFruit: String){
        self.ingFruit = ingFruit
    }
    func setIngWood(ingWood: String){
        self.ingWood = ingWood
    }
    func setIngSpecial(ingSpecial: String){
        self.ingSpecial = ingSpecial
    }
    func setPowerLevel(powerLevel: Int){
        self.powerLevel = powerLevel
    }
    
    func buildFragrance(name: String) -> Fragrance {
        return Fragrance(name: name, ingFlower: ingFlower, ingFruit: ingFruit, ingWood: ingWood, ingSpecial: ingSpecial, powerLevel: powerLevel)
    }
    
    class func getFragranceBuilder(fragranceType: FragranceType) -> FragranceBuilderFactory {
        switch fragranceType {
        case .Floral: return FloralFrangranceBuilder()
        case .Fruity: return FruityFrangranceBuilder()
        case .Oceanic: return OceanicFrangranceBuilder()
        }
    }
}

class FloralFrangranceBuilder: FragranceBuilderFactory {
    override init(){
        super.init()
        self.ingFlower = "Rose"
        self.powerLevel = 99
    }
    
}
class FruityFrangranceBuilder: FragranceBuilderFactory {
    override init(){
        super.init()
        self.ingFruit = "Apple"
        self.powerLevel = 99
    }
    
    
}
class OceanicFrangranceBuilder: FragranceBuilderFactory {
    override init(){
        super.init()
        self.ingWood = "Pine"
        self.powerLevel = 80
    }
}

let builder1 = FragranceBuilderFactory.getFragranceBuilder(fragranceType: FragranceType.Oceanic)
builder.setIngSpecial(ingSpecial: "Mermaid Scale")
let fragrance1 = builder.buildFragrance(name: "Ocean Song")
print(fragrance1.description)


//: 3. Use closure to config the FragranceBuilder
struct TearInHeavenFragrance {
    let name: String
    let ingFlower: String
    let ingFruit: String
    let ingWood: String
    let ingSpecial: String
    let powerLevel: Int
    var description: String {
        return "Fragrance:\(name)\n" +
            "Flower:\(ingFlower)\n" +
            "Fruit: \(ingFruit)\n" +
            "Wood: \(ingWood)\n" +
            "Special Ingredience: \(ingSpecial)\n" +
        "Power Level: \(powerLevel)\n"
    }
    
    init(name: String, builder: TearInHeavenFragranceBuilder) {
        self.name = name
        self.ingFlower = builder.ingFlower
        self.ingFruit = builder.ingFruit
        self.ingWood = builder.ingWood
        self.ingSpecial = builder.ingSpecial
        self.powerLevel = builder.powerLevel
    }
}

class TearInHeavenFragranceBuilder {
    var ingFlower = "Jasmine"
    var ingFruit = "Strawberry"
    var ingWood = "Sandalwood"
    var ingSpecial = "Fish's tear"
    var powerLevel = 0
    init(builderHandler: (TearInHeavenFragranceBuilder) -> Void) {
        builderHandler(self)
    }
}
let builder2 = TearInHeavenFragranceBuilder { (builder) -> Void in
    builder.ingFlower = "White Lily"
    builder.powerLevel = 78
}

let fragrance2 = TearInHeavenFragrance(name: "Tear in Heaven", builder: builder2)
print(fragrance2.description)




