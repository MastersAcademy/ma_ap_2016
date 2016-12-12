package base

type DataEntity interface {
	GetComplexData() map[string]string
}

type ComplexEntity struct {
	Id         int
	DataEntity DataEntity
	Formatter  BFormatter
}

func (ce *ComplexEntity) GetComplexData() map[string]string {
	return ce.DataEntity.GetComplexData()
}

func (ce *ComplexEntity) GetFormatted() string {
	return ce.Formatter.FormatData(*ce)
}

func (ce *ComplexEntity) SetFormatter(formatter BFormatter) {
	ce.Formatter = formatter
}

