package pool

import (
	"fmt"
	"errors"
	"sync"
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
