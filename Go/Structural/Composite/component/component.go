package component

type Component interface {
	AddChild(c Component)
	GetName() string
	GetChildren() []Component
	Print(prefix string) string
}
