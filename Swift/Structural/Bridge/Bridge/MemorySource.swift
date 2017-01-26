//
//  MemorySource.swift
//  Bridge
//
//  Created by Anton on 26.01.17.
//  Copyright © 2017 None. All rights reserved.
//

import Foundation

class MemorySource {
    var videos: [Video]
    var engine: ReportEngineProtocol
    
    init(videos: [Video], engine: ReportEngineProtocol) {
        self.videos = videos
        self.engine = engine
    }
}

extension MemorySource: SourceProtocol {
    func getVideos() throws -> String {
        return engine.buildReport(videos: self.videos)
    }
}
