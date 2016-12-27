package runeFlyweight

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	red     = string([]byte{27, 91, 57, 49, 59, 52, 48, 109})
	green   = string([]byte{27, 91, 57, 50, 59, 52, 48, 109})
	yellow  = string([]byte{27, 91, 57, 51, 59, 52, 48, 109})
	blue    = string([]byte{27, 91, 57, 52, 59, 52, 48, 109})
	magenta = string([]byte{27, 91, 57, 53, 59, 52, 48, 109})
	cyan    = string([]byte{27, 91, 57, 54, 59, 52, 48, 109})
	gray    = string([]byte{27, 91, 57, 55, 59, 52, 48, 109})
	reset   = string([]byte{27, 91, 48, 109})
	colors  = []string{red, green, yellow, blue, magenta, cyan, gray}
)

type Flyweight struct {
	r     rune
	color string
}

func (f *Flyweight) String() string {
	return fmt.Sprintf("%s%c%s", f.color, f.r, reset)
}

type Factory struct {
	sync.Mutex
	pool map[rune]*Flyweight
}

func NewFactory() *Factory {
	rand.Seed(time.Now().UTC().UnixNano())
	return &Factory{
		pool: map[rune]*Flyweight{},
	}
}

func (f *Factory) GetFlyweight(r rune) *Flyweight {
	f.Lock()
	defer f.Unlock()
	x, ok := f.pool[r]
	if !ok {
		x = &Flyweight{r, colors[rand.Intn(len(colors))]}
		f.pool[r] = x
	}
	return x
}

func (f *Factory) UniqueRunesNumber() int {
	return len(f.pool)
}
