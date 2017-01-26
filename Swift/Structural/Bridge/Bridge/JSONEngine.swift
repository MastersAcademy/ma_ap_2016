//
//  JSONEngine.swift
//  Bridge
//
//  Created by Anton on 26.01.17.
//  Copyright © 2017 None. All rights reserved.
//

import Foundation

class JSONEngine: ReportEngineProtocol {
    func buildReport(videos: [Video]) -> String {
        return Json.jsonString(from: videos)
    }
}
