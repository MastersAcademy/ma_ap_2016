//
//  SQLiteSource.swift
//  Bridge
//
//  Created by Anton on 26.01.17.
//  Copyright © 2017 None. All rights reserved.
//

import Foundation
import SQLite

class SQLiteSource {
    var engine: ReportEngineProtocol
    let db: Connection
    let videosTable = Table("videos")
    
    init?(dbPath: String, engine: ReportEngineProtocol) {
        guard let connection = try? Connection(dbPath) else {
    		return nil
    	}
        
        self.db = connection
        self.engine = engine
    }
    
    func createDB() {
        let idColumn = Expression<String>("id")
        let titleColumn = Expression<String>("title")
        let urlColumn = Expression<String>("url")
        
        do {
            try db.run(videosTable.create { t in
                t.column(idColumn)
                t.column(titleColumn)
                t.column(urlColumn)
            })

            
            try db.run(videosTable.insert(idColumn <- "1", titleColumn <- "Some funny video", urlColumn <- "https://youtube.com/funny"))
            try db.run(videosTable.insert(idColumn <- "2", titleColumn <- "Some ugly video", urlColumn <- "https://youtube.com/ugly"))

        } catch {
            print("Can't create db. Error: \(error)")
        }
    }
}

extension SQLiteSource: SourceProtocol {
    func getVideos() throws -> String {
        var videos: [Video] = []
        do {
            let idColumn = Expression<String>("id")
            let titleColumn = Expression<String>("title")
            let urlColumn = Expression<String>("url")

            for row in try db.prepare(videosTable) {
                let video = Video(id: row[idColumn],
                                  title: row[titleColumn],
                                  url: row[urlColumn])
            	videos.append(video)
            }
        } catch {
            throw error
        }
        
        return engine.buildReport(videos: videos)
    }
}

