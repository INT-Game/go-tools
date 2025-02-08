package log_context

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogValues_Set(t *testing.T) {
	v := &LogValues{}

	// Test normal set
	v.Set("key1", "value1")
	assert.Equal(t, []any{"key1", "value1"}, v.kvs)

	// Test update existing key
	v.Set("key1", "value2")
	assert.Equal(t, []any{"key1", "value2"}, v.kvs)

	// Test multiple sets
	v.Set("key2", "value2")
	assert.Equal(t, []any{"key1", "value2", "key2", "value2"}, v.kvs)
}

func TestLogValues_Get(t *testing.T) {
	// Test nil LogValues
	var nilV *LogValues
	val, ok := nilV.Get("any")
	assert.False(t, ok)
	assert.Nil(t, val)

	v := &LogValues{kvs: []any{"key1", "value1", "key2", 42}}

	// Test existing string value
	val1, ok := v.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", val1)

	// Test existing int value
	val2, ok := v.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, 42, val2)

	// Test non-existing key
	val3, ok := v.Get("key3")
	assert.False(t, ok)
	assert.Nil(t, val3)
}

func TestLogValues_All(t *testing.T) {
	// Test nil LogValues
	var v *LogValues
	assert.NotNil(t, v.All())
	assert.Empty(t, v.All())

	// Test non-nil LogValues
	v = &LogValues{kvs: []any{"key1", "value1"}}
	assert.Equal(t, []any{"key1", "value1"}, v.All())
}

func TestLogValues_GetStr(t *testing.T) {
	v := &LogValues{kvs: []any{"strKey", "string", "intKey", 42}}

	// Test successful string retrieval
	str1, ok := v.GetStr("strKey")
	assert.True(t, ok)
	assert.Equal(t, "string", str1)

	// Test non-string value
	str2, ok := v.GetStr("intKey")
	assert.False(t, ok)
	assert.Empty(t, str2)

	// Test non-existing key
	str3, ok := v.GetStr("nonexistent")
	assert.False(t, ok)
	assert.Empty(t, str3)
}

func TestLogValues_Copy(t *testing.T) {
	original := &LogValues{kvs: []any{"key1", "value1", "key2", 42}}

	// Test copy
	copied := original.Copy()
	assert.Equal(t, original.kvs, copied.kvs)

	// Verify deep copy by modifying original
	original.Set("key1", "modified")
	assert.NotEqual(t, original.kvs, copied.kvs)
}
