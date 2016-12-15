package main

import (
	"errors"
	"sync"
	"fmt"
	"time"
	"math/rand"
)

type ImageProcessor struct {
	id int
}

func (i *ImageProcessor)Process(data int) {
	fmt.Printf("Processor %v process %v\n", i.id, data)
}

func NewImageProcessor(id int) *ImageProcessor{
	return &ImageProcessor{
		id: id,
	}
}

var NoFreeProcessorError = errors.New("No Free processors")
var instance *ProcessorPool
var once sync.Once

type ProcessorPool struct {
	m sync.RWMutex
	idle map[int]*ImageProcessor
	used map[int]*ImageProcessor
	maxProcessors int
}

func GetProcessorPool() *ProcessorPool {
	once.Do(func() {
		instance = &ProcessorPool{
			idle: make(map[int]*ImageProcessor),
			used: make(map[int]*ImageProcessor),
			maxProcessors: 5,
		}
	})
	return instance
}

func (p *ProcessorPool)Acquire() (*ImageProcessor, error) {
	p.m.Lock()
	defer p.m.Unlock()
	if len(p.idle) == 0 && len(p.used) == p.maxProcessors {
		return nil, NoFreeProcessorError
	}

	if len(p.idle) > 0 {

		var processor *ImageProcessor
		var k int
		for k, processor = range p.idle {
			delete(p.idle, k)
			p.used[k] = processor
			return processor, nil
		}
	}
	id := len(p.idle) + len(p.used) + 1
	obj := NewImageProcessor(id)
	p.used[id] = obj
	return obj, nil

}

func (p *ProcessorPool)Release(obj *ImageProcessor) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.used, obj.id)
	p.idle[obj.id] = obj

}

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
					processor, _ := GetProcessorPool().Acquire()
					for processor == nil{
						time.Sleep(time.Duration(rand.Intn(10)))
						processor, _ = GetProcessorPool().Acquire()
					}
					processor.Process(w)
					GetProcessorPool().Release(processor)
					//time.Sleep(time.Duration(rand.Intn(10)))
				default:
					return
				}
			}
		}(i)
	}
	wg.Wait()
}