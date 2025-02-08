package perf

import (
	"context"
)

type MutexMapContext struct {
	BaseContext
}

func (c *MutexMapContext) Add(key string, value any) {
	c.Mutex.Lock()

	defer c.Mutex.Unlock()
	ctx := c.Ctx
	val := ctx.Value(LogContextKeyStr)
	kvs, ok := val.(*map[string]any)
	if !ok {
		kvs = &map[string]any{}
	}
	(*kvs)[key] = value
	c.Ctx = context.WithValue(ctx, LogContextKeyStr, kvs)
}

func (c *MutexMapContext) Get() any {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return c.Ctx.Value(LogContextKeyStr)
}
