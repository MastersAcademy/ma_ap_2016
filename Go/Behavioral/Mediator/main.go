package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/MastersAcademy/ma_ap_2016/Go/Behavioral/Mediator/messenger"
)

func main() {
	mm := messenger.NewMessageMediator()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT)

	for i := 0; i < 10; i++ {
		c := messenger.NewClient()
		c.SetMediator(mm)
		go func(c *messenger.Client, id int) {
			for {
				time.Sleep(100 * time.Millisecond)
				select {
				case msg := <-c.MsgChan:
					fmt.Printf("Client %d received message: '%s'\n", id, msg)
				default:
					if rand.Float64() > 0.999 {
						go c.SendMessage(fmt.Sprintf("Hello, I'm client %d", id))
					}
				}
			}
		}(c, i)
	}
	<-stopChan
}
