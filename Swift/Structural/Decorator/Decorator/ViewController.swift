//
//  ViewController.swift
//  Decorator
//
//  Created by Admin on 28.12.16.
//  Copyright © 2016 Admin. All rights reserved.
//

import UIKit

class ViewController: UIViewController {
    
    @IBOutlet weak var imageView: UIImageView!

    override func viewDidLoad() {
        super.viewDidLoad()

        imageView.imageWith(url: "https://lolkot.ru/wp-content/uploads/2013/07/o-privet_1374841583.jpg", completion: {
            print("loaded, image: \($0), error: \($1)")
        })
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }


}

