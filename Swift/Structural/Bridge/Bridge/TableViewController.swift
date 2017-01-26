//
//  TableViewController.swift
//  Bridge
//
//  Created by Anton on 26.01.17.
//  Copyright © 2017 None. All rights reserved.
//

import UIKit

enum SourceType {
    case memory
    case sqlite
}

enum ReportType {
    case csv
    case json
}

class TableViewController: UITableViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
	
        self.tableView.delegate = self
	}

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    override func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        let source: SourceType
        let report: ReportType
        
        switch indexPath.row {
        case 0:
            source = .sqlite
            report = .json
        case 1:
            source = .memory
            report = .json
        case 2:
            source = .memory
            report = .csv
        case 3:
            source = .sqlite
            report = .csv
        default:
            source = .memory
            report = .csv
        }
        
        if let viewController = storyboard!.instantiateViewController(withIdentifier: "ViewController") as? ViewController {
            viewController.reportType = report
            viewController.sourceType = source

            self.navigationController?.pushViewController(viewController, animated: true)
        }
        
    }
}
