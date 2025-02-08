package perf

import (
	"context"
	"sync"
)

type Key string

var LogContextKeyStr = Key("key")
var UnnamedKeyStr = "unamed_key"

type ContextInterface interface {
	Add(key string, value any)
	Get() any
	GetMutex() *sync.RWMutex
}

type BaseContext struct {
	ContextInterface
	Ctx   context.Context
	Mutex sync.RWMutex
}

func (c *BaseContext) GetMutex() *sync.RWMutex {
	return &c.Mutex
}
