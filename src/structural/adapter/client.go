package adapter

type Client struct {
	sender Sender
}

func NewClient(s Sender) *Client {
	return &Client{sender: s}
}

func (c *Client) Do() {
	_ = c.sender.Send("Hello from Client!")
}
