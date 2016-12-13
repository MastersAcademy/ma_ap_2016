import Foundation

class Singleton {
    // class constant with lazy initialization
    // also thread safe
    static let sharedInstance = Singleton()
    
    var number: Int = 0
    var description: String = ""
    
    private init() {}
    
    func printObj() {
        print("Singleton{number:\(self.number), description:\(self.description)}")
    }
}
