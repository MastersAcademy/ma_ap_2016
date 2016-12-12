package formatters

import "strconv"
import	"../base"

type JSONFormatter struct {
	keyPrefix string
}

func NewJSONFormatter(prefix string) *JSONFormatter {
	fmt := JSONFormatter{}
	fmt.keyPrefix = prefix
	return &fmt
}

func (jf *JSONFormatter) FormatData(complexEntity base.ComplexEntity) string {
	rmap := complexEntity.DataEntity.GetComplexData()
	result := "{ id:" + strconv.Itoa(complexEntity.Id) + "," + "data:{"
	for k := range rmap {
		result += jf.keyPrefix + k + ":" + rmap[k] + ","
	}
	result = result[0:len(result) - 1]
	result += "}"
	return result
}
