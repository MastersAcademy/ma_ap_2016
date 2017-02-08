package abstractMediator

type Mediator interface {
	AddChan(ch chan string)
	Send(msg string)
}

type Colleague interface {
	SetMediator(Mediator)
}
