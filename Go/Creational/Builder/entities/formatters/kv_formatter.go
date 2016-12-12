package formatters

import "strconv"
import "../base"

type KVFormatter struct {
	sep string
	cor string
}

func NewKVFormatter(sep string, cor string) *KVFormatter {
	fmt := KVFormatter{}
	if sep != "" {
		fmt.sep = sep
	} else {
		fmt.sep = "|"
	}
	if cor != "" {
		fmt.cor = cor
	} else {
		fmt.cor = "="
	}
	return &fmt
}

func (kvf *KVFormatter) FormatData(complexEntity base.ComplexEntity) string {
	rmap := complexEntity.DataEntity.GetComplexData()
	result := "id" + kvf.cor + strconv.Itoa(complexEntity.Id) + kvf.sep
	for k := range rmap {
		result += k + kvf.cor + rmap[k] + kvf.sep
	}
	result = result[0:len(result) - 1]
	return result
}

