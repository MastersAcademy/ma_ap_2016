package main

import (
	"sync"
	"time"
	"math/rand"
	"github.com/MastersAcademy/ma_ap_2016/Go/Creational/ObjectPool/pool"
)



func main(){
	workersCount := 10

	c := make(chan int, workersCount * 10)
	var wg sync.WaitGroup
	wg.Add(workersCount)
	for i:= 1; i<=workersCount*10; i++{
		c <- i
 	}

	for i:=0; i<workersCount; i++ {
		go func(id int){
			defer wg.Done()
			for {
				select{
				case w := <-c:
					processor, _ := pool.GetProcessorPool().Acquire()
					for processor == nil{
						time.Sleep(time.Duration(rand.Intn(10)))
						processor, _ = pool.GetProcessorPool().Acquire()
					}
					processor.Process(w)
					pool.GetProcessorPool().Release(processor)
				default:
					return
				}
			}
		}(i)
	}
	wg.Wait()
}