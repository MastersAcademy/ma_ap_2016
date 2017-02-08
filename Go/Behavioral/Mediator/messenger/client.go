package messenger

import "github.com/MastersAcademy/ma_ap_2016/Go/Behavioral/Mediator/abstractMediator"

type Client struct {
	mediator abstractMediator.Mediator
	MsgChan  chan string
}

func NewClient() *Client {
	return &Client{
		MsgChan: make(chan string, 1),
	}
}

func (c *Client) SetMediator(m abstractMediator.Mediator) {
	m.AddChan(c.MsgChan)
	c.mediator = m
}

func (c *Client) SendMessage(msg string) {
	c.mediator.Send(msg)
}
