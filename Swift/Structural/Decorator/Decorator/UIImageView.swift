//
//  UIImageView.swift
//  Decorator
//
//  Created by Admin on 28.12.16.
//  Copyright © 2016 Admin. All rights reserved.
//

import Foundation
import UIKit

enum UIImageLoaderError: Error {
    case wrongUrl
    case dataFromUrlIsNil
    case cantBuildImageFromData
}

extension UIImageView {
    typealias ImageWithUrlCompletion = (UIImage?, Error?)->Void
    func imageWith(url: String, completion: ImageWithUrlCompletion?) {
        guard let url = URL(string: url) else {
            if let completion = completion {
                completion(nil, UIImageLoaderError.wrongUrl)
            }
            return
        }
        
        URLSession.shared.dataTask(with: url) {
            (data, response, error) in
            
            guard error == nil else {
                if let completion = completion {
                    completion(nil, error)
                }
                return
            }
            
            guard let data = data else {
                if let completion = completion {
                    completion(nil, UIImageLoaderError.dataFromUrlIsNil)
                }
                return
            }
            
            guard let image = UIImage(data: data) else {
                if let completion = completion {
                    completion(nil, UIImageLoaderError.cantBuildImageFromData)
                }
                return
            }
            
            DispatchQueue.main.async {
                [weak self] () in
                self?.image = image
                if let completion = completion {
                    completion(image, nil)
                }
            }
            
        }.resume()
    }
}
