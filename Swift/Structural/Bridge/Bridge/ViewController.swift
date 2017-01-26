//
//  ViewController.swift
//  Bridge
//
//  Created by Anton on 26.01.17.
//  Copyright © 2017 None. All rights reserved.
//

import UIKit

class ViewController: UIViewController {
    
    @IBOutlet weak var textView: UITextView!
    
    var sourceType: SourceType?
    var reportType: ReportType?
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        guard let source = createSourceFrom(sourceType: sourceType ?? .memory, reportType: reportType ?? .json) else {
            print("Can't create source")
            return
        }
        
        do {
        	textView.text = try source.getVideos()
        } catch {
            print("Can't get videos. Error: \(error)")
        }
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        
    }
    
    func createSourceFrom(sourceType: SourceType, reportType: ReportType) -> SourceProtocol? {
        let engine: ReportEngineProtocol
        switch reportType {
        case .csv:
            engine = CSVEngine()
        case .json:
            engine = JSONEngine()
        }
        
        let source: SourceProtocol?
        switch sourceType {
        case .memory:
            let video1 = Video(id: "1", title: "Video from memory 1", url: "https://url.com/1")
            let video2 = Video(id: "2", title: "Video from memory 2", url: "https://url.com/2")
            
            source = MemorySource(videos: [video1, video2], engine: engine)
        case .sqlite:
            let path = NSSearchPathForDirectoriesInDomains(.documentDirectory, .userDomainMask, true ).first!
            source = SQLiteSource(dbPath: "\(path)/my.db", engine: engine)
        }
        
        return source
    }
}

