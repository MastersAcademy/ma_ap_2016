package ReportEngines

import "github.com/MastersAcademy/ma_ap_2016/Go/Structural/Bridge/Models"

type ReportEngine interface {
	BuildReport(videos []Models.Video) string
}
