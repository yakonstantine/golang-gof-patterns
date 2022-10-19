package adapter

type OldSenderAdapter struct {
	snd *OldSender
}

func NewOldSenderAdapter(os *OldSender) *OldSenderAdapter {
	return &OldSenderAdapter{snd: os}
}

func (a *OldSenderAdapter) Send(msg string) error {
	a.snd.SendMessage(msg)
	return nil
}
