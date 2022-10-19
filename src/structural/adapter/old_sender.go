package adapter

import "fmt"

type OldSender struct {
}

func (s *OldSender) SendMessage(msg string) {
	fmt.Println("OldSender sent: '", msg, "'")
}
