package component

import (
	"bytes"
)

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func (f *File) AddChild(c Component) {
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetChildren() []Component {
	return []Component{}
}

func (f *File) Print(prefix string) string {
	var b bytes.Buffer
	b.WriteString(prefix)
	b.WriteRune('/')
	b.WriteString(f.name)
	b.WriteRune('\n')
	return b.String()
}
