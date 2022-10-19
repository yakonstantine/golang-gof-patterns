package composite

type FSComponent interface {
	Name() string
	Search(text string) (string, bool)
}
