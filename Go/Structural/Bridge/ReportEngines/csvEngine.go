package ReportEngines

import (
	"github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/Models"
	"bytes"
	"fmt"
)

type CSVEngine struct {

}

func NewCSVEngine() ReportEngine {
	return &CSVEngine{}
}

func (e *CSVEngine) BuildReport(videos []Models.Video) string  {
	var buffer bytes.Buffer
	buffer.WriteString("Id, Title, Url\n")

	for _, v := range videos {
		buffer.WriteString(fmt.Sprintf("%v, '%s', %s\n", v.ID, v.Title, v.URL))
	}

	return buffer.String()
}
