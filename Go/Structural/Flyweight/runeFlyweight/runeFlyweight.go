package runeFlyweight

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	red     = "\x1b[;91m"
	green   = "\x1b[;92m"
	yellow  = "\x1b[;93m"
	blue    = "\x1b[;94m"
	magenta = "\x1b[;95m"
	cyan    = "\x1b[;96m"
	white   = "\x1b[;97m"
	reset   = "\x1b[0m"

	colors = []string{red, green, yellow, blue, magenta, cyan, white}
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
