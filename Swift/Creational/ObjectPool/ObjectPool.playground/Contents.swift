/*:
 Master Academy - Design Patterns implemented in Swift 3.0
 */
/*:
 # ObjectPool
 The object pool pattern is a software creational design pattern that uses a set of initialized objects kept ready to use – a "pool" – rather than allocating and destroying them on demand. A client of the pool will request an object from the pool and perform operations on the returned object. When the client has finished, it returns the object to the pool rather than destroying it; this can be done manually or automatically.
 
 Object pools are primarily used for performance: in some circumstances, object pools significantly improve performance. Object pools complicate object lifetime, as objects obtained from and returned to a pool are not actually created or destroyed at this time, and thus require care in implementation.
 */

import UIKit
import Foundation

struct MagicObject {
    let name: String
    let serialNumber: Int
    var occupier: [String] = []
    var borrowedCount: Int = 0
}

//: a basic pool which allows getting and returning any objects.
class Pool<T> {
    public private(set) var data = [T]()
    
    init(items: [T]) {
        data.reserveCapacity(data.count)
        for item in items {
            data.append(item)
        }
    }
    
    func getFromPool() -> T? {
        var result: T?
            if data.count > 0 {
                result = self.data.removeFirst()
            }
        return result
    }
    
    func returnToPool(item: T){
        self.data.append(item)
    }
}

extension Int {
    func times(action: (Int)->()) {
        for i in 0..<self {
            action(i)
        }
    }
}

class MagicHouse {
    private let pool: Pool<MagicObject>
    static var sharedInstance = MagicHouse()
    static var magicDebtInfo:[(String, Int, String)] = []
    private init(){
        var magicObjects:[MagicObject] = []
        2.times{
            magicObjects.append(MagicObject(name: "Red Diamond", serialNumber: $0, occupier: [], borrowedCount: 0))
        }
        3.times{
            magicObjects.append(MagicObject(name: "Blue Heart", serialNumber: $0, occupier: [], borrowedCount: 0))
        }
        self.pool = Pool(items: magicObjects)
    }
    
    class func lendMagicObject(occupier: String) -> MagicObject? {
        var magicObject = sharedInstance.pool.getFromPool()
        if magicObject != nil {
            magicObject!.occupier.append(occupier)
            magicObject!.borrowedCount += 1
            magicDebtInfo.append((magicObject!.name, magicObject!.serialNumber, occupier))
            print("\(occupier) is borrowing \(magicObject!.name) #\(magicObject!.serialNumber)")
        }
        return magicObject
    }
    
    class func receiveMagicObject(obj: MagicObject) {
        magicDebtInfo = magicDebtInfo.filter{
            $0.0 != obj.name && $0.1 != obj.serialNumber
        }
        sharedInstance.pool.returnToPool(item: obj)
        print("\(obj.occupier.last!) returning \(obj.name) #\(obj.serialNumber)")
    }
    
    class func printReport(){
        print("\nShow Report: Magic House currently has \(sharedInstance.pool.data.count) magic object(s) in stock")
        (sharedInstance.pool.data as [MagicObject]).map {
            print("\($0.name) #\($0.serialNumber) \nBorrowed \($0.borrowedCount) time(s) by \($0.occupier)")
        }
        
        if magicDebtInfo.count > 0 {
            print("\nMagic Objects currently lent out:")
            magicDebtInfo.map{
                print("\($0.0) #\($0.1) by \($0.2)")
            }
        }
    }
}

//: Testing
print("\n------Starting test...")

for i in 1 ... 7 {
    var obj = MagicHouse.lendMagicObject(occupier: "person #\(i)")
    if obj != nil {
        Thread.sleep(forTimeInterval: Double(arc4random_uniform(3)))
        MagicHouse.receiveMagicObject(obj: obj!)
    }
}

let m1 = MagicHouse.lendMagicObject(occupier: "William")
let m3 = MagicHouse.lendMagicObject(occupier: "Tato")
MagicHouse.printReport()



