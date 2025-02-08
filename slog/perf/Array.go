package perf

import "context"

type ArrayContext struct {
	BaseContext
}

func (c *ArrayContext) Add(key string, value any) {
	ctx := c.Ctx
	val := ctx.Value(LogContextKeyStr)
	kvs, ok := val.([]any)
	if !ok {
		kvs = []any{}
		if val != nil {
			kvs = append(kvs, UnnamedKeyStr, val)
		}
	}
	has := false
	for i, k := range kvs {
		if k == key && i+1 < len(kvs) {
			kvs[i+1] = value
			has = true
			break
		}
	}
	if !has {
		kvs = append(kvs, key, value)
	}
	c.Ctx = context.WithValue(ctx, LogContextKeyStr, kvs)
}

func (c *ArrayContext) Get() any {
	return c.Ctx.Value(LogContextKeyStr)
}
