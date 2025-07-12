package unit00base

import (
	"sync"
	"time"

	"github.com/u00io/gomisc/logger"
)

type Unit struct {
	mtx     sync.Mutex
	started bool
	stoping bool
	config  map[string]string
	values  map[string]string
	iUnit   IUnit
}

type IUnit interface {
	Start()
	SetConfig(config map[string]string)
	GetValue(key string) string
	SetValue(key, value string)
	Stop()
	Tick()
}

func (c *Unit) Init(iUnit IUnit) {
	c.config = make(map[string]string)
	c.values = make(map[string]string)
	c.iUnit = iUnit
}

func (c *Unit) SetConfig(config map[string]string) {
	c.mtx.Lock()
	c.config = config
	c.mtx.Unlock()
}

func (c *Unit) GetValue(key string) string {
	c.mtx.Lock()
	value, exists := c.values[key]
	c.mtx.Unlock()
	if !exists {
		return ""
	}
	return value
}

func (c *Unit) SetValue(key, value string) {
	c.mtx.Lock()
	c.values[key] = value
	c.mtx.Unlock()
}

func (c *Unit) Start() {
	c.mtx.Lock()
	if c.started {
		c.mtx.Unlock()
		return
	}
	c.stoping = false
	go c.thWork()
	c.mtx.Unlock()
}

func (c *Unit) Stop() {
	c.mtx.Lock()
	if !c.started {
		c.mtx.Unlock()
		return
	}
	c.stoping = true
	c.mtx.Unlock()

	dtStartWaitingForStop := time.Now()

	for {
		if time.Since(dtStartWaitingForStop) > 1*time.Second {
			logger.Println("Unit stop timeout exceeded, force stopping")
			break
		}
		c.mtx.Lock()
		started := c.started
		c.mtx.Unlock()
		if !started {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func (c *Unit) thWork() {
	c.started = true
	for !c.stoping {
		c.iUnit.Tick()
		time.Sleep(100 * time.Millisecond)
	}
	c.started = false
	c.stoping = false
}

func (c *Unit) Tick() {
}
