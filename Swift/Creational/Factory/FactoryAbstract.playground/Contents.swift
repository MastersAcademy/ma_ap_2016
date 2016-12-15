//: Playground - noun: a place where people can play

import UIKit
import Foundation

protocol Dough {}
protocol Sauce {}
protocol Cheese {}
protocol Veggies {}
protocol Pepperoni {}
protocol Clams {}

protocol PizzaIngredientFactory {
    func createDough() -> Dough
    func createSauce() -> Sauce
    func createCheese() -> Cheese
    func createVeggies() -> [Veggies]
    func createPepperoni() -> Pepperoni
    func createClam() -> Clams
}

class NYPizzaIngredientFactory: PizzaIngredientFactory {
    func createDough() -> Dough {
        return ThinCrustDough()
    }
    
    func createSauce() -> Sauce {
        return MarinaraSauce()
    }
    
    func createCheese() -> Cheese {
        return ReggianoCheese()
    }
    
    func createVeggies() -> [Veggies] {
        let veggies: [Veggies] = [Garlic(), Onion(), Mushroom(), RedPepper()]
        return veggies
    }
    
    func createPepperoni() -> Pepperoni {
        return SlicedPepperoni()
    }
    
    func createClam() -> Clams {
        return FreshClams()
    }
}

class ChikagoPizzaIngredientFactory: PizzaIngredientFactory {
    func createDough() -> Dough {
        return ThinCrustDough()
    }
    
    func createSauce() -> Sauce {
        return MarinaraSauce()
    }
    
    func createCheese() -> Cheese {
        return ReggianoCheese()
    }
    
    func createVeggies() -> [Veggies] {
        let veggies: [Veggies] = [Garlic(), Onion(), Mushroom(), RedPepper()]
        return veggies
    }
    
    func createPepperoni() -> Pepperoni {
        return SlicedPepperoni()
    }
    
    func createClam() -> Clams {
        return FrozenClams()
    }
}

class Garlic: Veggies {}
class Onion: Veggies {}
class Mushroom: Veggies {}
class RedPepper: Veggies {}
class ThinCrustDough: Dough {}
class MarinaraSauce: Sauce {}
class ReggianoCheese: Cheese {}
class SlicedPepperoni: Pepperoni {}
class FreshClams: Clams {}
class FrozenClams: Clams {}


protocol Pizza {
    var name: String! { get set }
    var dough: Dough! { get }
    var sauce: Sauce! { get }
    var veggies: [Veggies]? { get }
    var cheese: Cheese? { get }
    var pepperoni: Pepperoni? { get }
    var clam: Clams? { get }
    
    // abstract method
    func prepare()
    
    func bake()
    func cut()
    func box()
}

extension Pizza {
    func bake() {
        print("Bake for 25 minutes at 350")
    }
    func cut() {
        print("Cutting the pizza into diagonal slices")
    }
    func box() {
        print("Place pizza in official PizzaStore box")
    }
}

class CheesePizza: Pizza {
    var name: String!
    var dough: Dough!
    var sauce: Sauce!
    var veggies: [Veggies]?
    var cheese: Cheese?
    var pepperoni: Pepperoni?
    var clam: Clams?
    
    var ingredientFactory: PizzaIngredientFactory
    
    init(ingredientFactory: PizzaIngredientFactory) {
        self.ingredientFactory = ingredientFactory
    }
    
    func prepare() {
        print("Preparing \(name)")
        dough = ingredientFactory.createDough()
        sauce = ingredientFactory.createSauce()
        cheese = ingredientFactory.createCheese()
    }
}

class ClamPizza: Pizza {
    var name: String!
    var dough: Dough!
    var sauce: Sauce!
    var veggies: [Veggies]?
    var cheese: Cheese?
    var pepperoni: Pepperoni?
    var clam: Clams?
    
    var ingredientFactory: PizzaIngredientFactory
    
    init(ingredientFactory: PizzaIngredientFactory) {
        self.ingredientFactory = ingredientFactory
    }
    
    func prepare() {
        print("Preparing \(name)")
        dough = ingredientFactory.createDough()
        sauce = ingredientFactory.createSauce()
        cheese = ingredientFactory.createCheese()
        clam = ingredientFactory.createClam()
    }
}

protocol PizzaStore {
    func orderPizza(type: String) -> Pizza
    func createPizza(type: String) -> Pizza
}

extension PizzaStore {
    func orderPizza(type: String) -> Pizza {
        let pizza = createPizza(type: type)
        
        pizza.prepare()
        pizza.bake()
        pizza.cut()
        pizza.box()
        
        return pizza
    }
}

class NYPizzaStore: PizzaStore {
    
    func createPizza(type: String) -> Pizza {
        let ingredientFactory = NYPizzaIngredientFactory()
        let newPizza: Pizza
        
        switch type {
        case "cheese":
            newPizza = CheesePizza(ingredientFactory: ingredientFactory)
            newPizza.name = "New York Style Cheese Pizza"
        case "clam":
            newPizza = ClamPizza(ingredientFactory: ingredientFactory)
            newPizza.name = "New York Style Clam Pizza"
        default:
            newPizza = CheesePizza(ingredientFactory: ingredientFactory)
            newPizza.name = "New York Style Cheese Pizza"
        }
        
        return newPizza
    }
}

class ChikagoPizzaStore: PizzaStore {
    // fabric method
    func createPizza(type: String) -> Pizza {
        let ingredientFactory: PizzaIngredientFactory = ChikagoPizzaIngredientFactory()
        let newPizza: Pizza
        
        switch type {
        case "cheese":
            newPizza = CheesePizza(ingredientFactory: ingredientFactory)
            newPizza.name = "Chikago Style Cheese Pizza"
        case "clam":
            newPizza = ClamPizza(ingredientFactory: ingredientFactory)
            newPizza.name = "Chikago Style Clam Pizza"
        default:
            newPizza = CheesePizza(ingredientFactory: ingredientFactory)
            newPizza.name = "Chikago Style Cheese Pizza"
        }
        return newPizza
    }
}


let nyStore = NYPizzaStore()
let chikagoStore = ChikagoPizzaStore()
var pizza: Pizza
pizza = nyStore.orderPizza(type: "cheese")
print(pizza.name)
print("\n")
pizza = chikagoStore.orderPizza(type: "cheese")
print(pizza.name)
