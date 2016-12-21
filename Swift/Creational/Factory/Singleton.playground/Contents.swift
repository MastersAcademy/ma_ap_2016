/*:
 Master Academy - Design Patterns implemented in Swift 3.0
 */
/*:
 # Singleton
 “The lazy initializer for a global variable (also for static members of structs and enums)
 is run the first time that global is accessed,
 and is launched as `dispatch_once` to make sure that the initialization is atomic.
 This enables a cool way to use `dispatch_once` in your code:
 just declare a global variable with an initializer and mark it private.”
 — Apple's Swift Blog
 */

import UIKit
import Foundation

class SomeClass {
    static let sharedInstance = SomeClass()
    
    var number: Int = 0
    var someDescription: String = ""
    
    private init() { }
    
    var description: String {
        return "Singleton{number:\(self.number), someDescription:\(self.someDescription)}"
    }
    
    // other method
}


let x = SomeClass.sharedInstance
x.number = 5
print("x = \(x.description)")

let y = SomeClass.sharedInstance
y.someDescription = "test"
print("x = \(x.description)")
print("y = \(y.description)")