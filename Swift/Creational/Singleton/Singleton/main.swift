import Foundation

let x = Singleton.sharedInstance
x.number = 5
print(x)

let y = Singleton.sharedInstance
y.someDescription = "test"
print(x)
print(y)


