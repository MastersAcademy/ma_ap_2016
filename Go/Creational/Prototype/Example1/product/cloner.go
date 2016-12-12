package product

type Cloner interface {
	Clone() Cloner
	GetDescription() string
	SetDescription(string)
}
