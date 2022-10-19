package composite

import "strings"

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{name: name}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Search(text string) (string, bool) {
	if strings.Contains(f.name, text) {
		return f.name, true
	}
	return "", false
}
