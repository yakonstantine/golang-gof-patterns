package composite

import "strings"

type Folder struct {
	components []FSComponent
	name       string
}

func NewFolder(name string) *Folder {
	return &Folder{name: name}
}

func (f *Folder) Name() string {
	return f.name
}

func (f *Folder) Add(cmp FSComponent) {
	f.components = append(f.components, cmp)
}

func (f *Folder) Search(text string) (string, bool) {
	if strings.Contains(f.name, text) {
		return f.name, true
	}
	for _, cmp := range f.components {
		if nm, ok := cmp.Search(text); ok {
			return nm, ok
		}
	}
	return "", false
}
