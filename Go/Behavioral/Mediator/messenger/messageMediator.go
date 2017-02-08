package messenger

type MessageMediator struct {
	msgChans [](chan string)
}

func NewMessageMediator() *MessageMediator {
	return &MessageMediator{}
}

func (m *MessageMediator) AddChan(ch chan string) {
	m.msgChans = append(m.msgChans, ch)
}

func (m *MessageMediator) Send(msg string) {
	for i := 0; i < len(m.msgChans); i++ {
		m.msgChans[i] <- msg
	}
}
