package adapter

import "fmt"

type EmailSender struct {
}

func (es *EmailSender) Send(msg string) error {
	fmt.Println("EmailSender sent: '", msg, "'")
	return nil
}
