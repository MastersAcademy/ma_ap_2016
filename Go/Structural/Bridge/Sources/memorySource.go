package Sources

import (
	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/Models"
	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/ReportEngines"
)

type MemorySource struct {
	videos []Models.Video
	engine ReportEngines.ReportEngine
}

func NewMemorySource(videos []Models.Video, engine ReportEngines.ReportEngine) Source {
	return &MemorySource{
		videos: videos,
		engine: engine,
	}
}

func (s *MemorySource) GetVideos() (string, error) {
	return s.engine.BuildReport(s.videos), nil
}
