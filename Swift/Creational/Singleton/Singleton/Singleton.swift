import Foundation

class Singleton: CustomStringConvertible {
    // class constant with lazy initialization
    // also thread safe
    static let sharedInstance = Singleton()
    
    var number: Int = 0
    var someDescription: String = ""
    
    private init() {}
    
    var description: String {
        return "Singleton{number:\(self.number), someDescription:\(self.someDescription)}"
    }
}
