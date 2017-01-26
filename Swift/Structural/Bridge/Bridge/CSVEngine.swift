//
//  CSVEngine.swift
//  Bridge
//
//  Created by Anton on 26.01.17.
//  Copyright © 2017 None. All rights reserved.
//

import Foundation

class CSVEngine : ReportEngineProtocol {
    func buildReport(videos: [Video]) -> String {
        var output = "Id, Title, Url\n"
        
        for video in videos {
            output += String(format: "%@, '%@', %@\n", video.id, video.title, video.url)
        }
        
        return output
    }
}
