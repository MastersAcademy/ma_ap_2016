//
//  ReportEngineProtocol.swift
//  Bridge
//
//  Created by Anton on 26.01.17.
//  Copyright © 2017 None. All rights reserved.
//

import Foundation

protocol ReportEngineProtocol {
    func buildReport(videos: [Video]) -> String
}
