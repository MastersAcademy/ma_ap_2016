//
//  Json.swift
//  Reflection
//
//  Created by Anton on 18.01.17.
//  Copyright © 2017 None. All rights reserved.
//

import Foundation

protocol Numeric { }

extension Int : Numeric {}
extension Float : Numeric {}
extension Double : Numeric {}
extension UInt : Numeric {}
extension Bool : Numeric {}

open class Json {
    open class func jsonString(from object: Any, depth: UInt = 0) -> String {
        let mirror = Mirror(reflecting: object)
        let prefix = prefixFor(depth: depth)

        return "{\n\(jsonString(from: mirror, depth: depth + 1))\n\(prefix)}"
    }
    
    fileprivate class func jsonString(from object: Mirror, depth: UInt) -> String {
        let prefix = prefixFor(depth: depth)
        
        var outputString = ""
        var isFirstItr = true
        
        // iterate through objects properties
        for case let (label?, value) in object.children {
            if !isFirstItr {
                outputString += ",\n"
            }
            
            outputString += "\(prefix)\"\(label)\" : " + generateJsonStringFor(value: value, depth: depth)

            isFirstItr = false
        }
        
        return outputString
    }
    
    fileprivate class func generateJsonStringFor(value: Any, depth: UInt) -> String {
        let value = unwrap(value)
        
        switch value {
        case is Numeric:
            return "\(value)"
        case is String:
            return "\"\(value)\""
        case is Array<Any>:
            let prefix = prefixFor(depth: depth + 1)
            var result = "[\n"
            var isFirstItr = true

            for valueItr in value as! Array<Any> {
                if !isFirstItr {
                    result += ",\n"
                }
                result += prefix + generateJsonStringFor(value: valueItr, depth: depth + 1)
                isFirstItr = false
            }
            result += "\n" + prefixFor(depth: depth) + "]"
            return result
        case is Dictionary<AnyHashable, Any>:
            let prefix = prefixFor(depth: depth + 1)
            var result = "{\n"
            var isFirstItr = true
            
            for (key, valueItr) in value as! Dictionary<AnyHashable, Any> {
                if !isFirstItr {
                    result += ",\n"
                }
                result += prefix + "\"\(key)\" : " + generateJsonStringFor(value: valueItr, depth: depth + 1)
                isFirstItr = false
            }
            result += "\n" + prefixFor(depth: depth) + "}"
            return result
        default:
            return jsonString(from: value, depth: depth)
        }
    }
    
    fileprivate class func prefixFor(depth: UInt) -> String {
        guard depth > 0 else {
            return ""
        }
        
        var prefix = ""
        for _ in 1...depth {
            prefix += "\t"
        }
        
        return prefix
    }
    
    fileprivate class func unwrap(_ any: Any) -> Any {
        let mi = Mirror(reflecting: any)
        if mi.displayStyle != .optional {
            return any
        }
        
        if mi.children.count == 0 { return NSNull() }
        let (_, some) = mi.children.first!
        return some
    }
}

