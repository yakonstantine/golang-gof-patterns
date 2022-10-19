package adapter

type Sender interface {
	Send(msg string) error
}
