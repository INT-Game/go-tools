package perf

import (
	"context"
	"sync"
)

type SyncMapContext struct {
	BaseContext
}

func (c *SyncMapContext) Add(key string, value any) {
	val := c.Ctx.Value(LogContextKeyStr)
	kvs, ok := val.(*sync.Map)
	if !ok {
		kvs = &sync.Map{}
		if val != nil {
			kvs.Store(UnnamedKeyStr, val)
		}
	}
	kvs.Store(key, value)
	c.Ctx = context.WithValue(c.Ctx, LogContextKeyStr, kvs)
}

// The log of SyncMap doesn't output as expected, duo to zap doesn't support sync.Map directly.
func (c *SyncMapContext) Get() any {
	return c.Ctx.Value(LogContextKeyStr)
}
