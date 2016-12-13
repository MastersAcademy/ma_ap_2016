import Foundation

let x = Singleton.sharedInstance
x.number = 5
x.printObj()

let y = Singleton.sharedInstance
y.description = "test"
x.printObj()
y.printObj()

