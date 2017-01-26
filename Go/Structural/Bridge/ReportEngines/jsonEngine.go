package ReportEngines

import (
	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/Models"
	"encoding/json"
	"log"
)

type JsonEngine struct {

}

func NewJsonEngine() ReportEngine {
	return &JsonEngine{}
}

func (e *JsonEngine) BuildReport(videos []Models.Video) string  {
	b, err := json.Marshal(videos)
	if err != nil {
		log.Panic("can't serialize json")
	}
	return string(b)
}
