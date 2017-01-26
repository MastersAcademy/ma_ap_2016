package main

import (
	"os"
	"fmt"
	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/Sources"
	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/ReportEngines"
	"log"
	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/Models"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf("You must specify source (db|memory) and report engine (json|csv)")
	}

	var source Sources.Source
	var engine ReportEngines.ReportEngine
	var err error

	switch args[1] {
	case "json":
		engine = ReportEngines.NewJsonEngine()
	case "csv":
		engine = ReportEngines.NewCSVEngine()
	default:
		log.Panic("Wrong engine")
	}

	switch args[0] {
	case "db":
		if len(args) < 3 {
			log.Panic("You must provide connection string for db source")
		}
		source, err = Sources.NewPostgresSource(args[2], engine)
		if err != nil {
			log.Panic(err)
		}
	case "memory":
		source = Sources.NewMemorySource(
			[]Models.Video{
				Models.Video{
					ID: "first",
					Title: "First video",
					URL: "http://some.com/first.mp4",
				},
				Models.Video{
					ID: "second",
					Title: "Second video",
					URL: "http://some.com/second.mp4",
				},
			},
			engine,
		)
	}

	report, err := source.GetVideos()
	if err != nil {
		log.Panic(err)
	}
	println(report)
}
