package log_context

type LogValues struct {
	kvs []any
}

func (v *LogValues) Set(key string, value any) {
	has := false
	for i, k := range v.kvs {
		if k == key && i+1 < len(v.kvs) {
			v.kvs[i+1] = value
			has = true
			break
		}
	}
	if !has && len(v.kvs) < ArrayLimit {
		v.kvs = append(v.kvs, key, value)
	}
}

func (v *LogValues) Get(key string) (value any, ok bool) {
	for i, k := range v.All() {
		if k == key && i+1 < len(v.kvs) {
			value = v.kvs[i+1]
			return value, true
		}
	}
	return nil, false
}

func (v *LogValues) All() []any {
	if v == nil {
		return []any{}
	}
	return v.kvs
}

func (v *LogValues) GetStr(key string) (value string, ok bool) {
	val, ok := v.Get(key)
	if !ok {
		return "", ok
	}
	value, ok = val.(string)
	return value, ok
}

func (v *LogValues) Copy() *LogValues {
	lv := &LogValues{}
	lv.kvs = make([]any, len(v.kvs))
	copy(lv.kvs, v.kvs)
	return lv
}
