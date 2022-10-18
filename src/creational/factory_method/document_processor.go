package factorymethod

type DocumentProcessor interface {
	Write(str string)
	Render()
}
