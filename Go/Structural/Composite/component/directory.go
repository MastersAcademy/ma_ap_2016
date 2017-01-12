package component

import (
	"bytes"
)

type Directory struct {
	name     string
	children []Component
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name: name,
	}
}

func (d *Directory) AddChild(c Component) {
	d.children = append(d.children, c)
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetChildren() []Component {
	return d.children
}

func (d *Directory) Print(prefix string) string {
	var p bytes.Buffer
	p.WriteString(prefix)
	p.WriteRune('/')
	p.WriteString(d.name)
	b := p
	b.WriteRune('\n')
	for _, c := range d.children {
		b.WriteString(c.Print(p.String()))
	}
	return b.String()
}
