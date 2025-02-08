package slog

import (
	"context"
	"fmt"
	"os"
	"slog/loggers"
	"slog/perf"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestPerformance(t *testing.T) {
	eleCount := 100
	count := 1
	runSinglePerformance(&perf.ArrayContext{BaseContext: perf.BaseContext{
		Ctx: context.Background(),
	}}, count, "Array")
	runSinglePerformance(&perf.MutexMapContext{BaseContext: perf.BaseContext{
		Ctx: context.Background(),
	}}, count, "MutexMap")
	runSinglePerformance(&perf.SyncMapContext{BaseContext: perf.BaseContext{
		Ctx: context.Background(),
	}}, count, "SyncMap")

	runConcurrentTest(&perf.ArrayContext{BaseContext: perf.BaseContext{
		Ctx: context.Background(),
	}}, count, "Array", eleCount)
	runConcurrentTest(&perf.MutexMapContext{
		BaseContext: perf.BaseContext{
			Ctx: context.Background(),
		}}, count, "MutexMap", eleCount)
	runConcurrentTest(&perf.SyncMapContext{BaseContext: perf.BaseContext{
		Ctx: context.Background(),
	}}, count, "SyncMap", eleCount)

	err := os.RemoveAll("./tmp")
	if err != nil {
		t.Errorf("Failed to remove temporary directory: %v", err)
	}
}
func runSinglePerformance(ctx perf.ContextInterface, count int, name string) {
	loggers.DefaultPrintf("runSinglePerformance[%s] ", name)
	Init(LogConfig{
		Dir:  fmt.Sprintf("./tmp/single/%s", name),
		File: true,
		RotateConfig: &RotateConfig{
			MaxSize:    1000,
			MaxAge:     10,
			MaxBackups: 10,
			Compress:   false,
		},
	})
	cur := time.Now()
	loggers.DefaultPrintln("\tstart time: ", cur)
	for i := 0; i < count; i++ {
		ctx.Add("key"+strconv.Itoa(i%5), "value"+strconv.Itoa(i%5))
		result := ctx.Get()
		CInfow(context.TODO(), "test", "ctx", result)
	}
	end := time.Now()
	loggers.DefaultPrintln("\tend time: ", end)
	loggers.DefaultPrintln("\tduration: ", end.Sub(cur))
	Close()
}

func runConcurrentTest(ctx perf.ContextInterface, count int, name string, eleCount int) {
	loggers.DefaultPrintf("runConcurrentTest[%s] ", name)
	Init(LogConfig{
		Dir:  fmt.Sprintf("./tmp/concurrent/%s", name),
		File: true,
		RotateConfig: &RotateConfig{
			MaxSize:    1000,
			MaxAge:     10,
			MaxBackups: 10,
			Compress:   false,
		},
	})
	var wg sync.WaitGroup
	cur := time.Now()
	loggers.DefaultPrintln("\tstart time: ", cur)
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			ctx.Add("key"+strconv.Itoa(i%eleCount), "value"+strconv.Itoa(i%eleCount))
			result := ctx.Get()
			array, ok := result.([]any)
			if ok {
				CInfow(context.TODO(), "array", array...)
			} else {
				mutexMap, ok := result.(*map[string]interface{})
				if ok {
					ctx.GetMutex().RLock()
					CInfow(context.TODO(), "mutexMap", "ctx", mutexMap)
					ctx.GetMutex().RUnlock()
				} else {
					CInfow(context.TODO(), "other", "ctx", result)
				}
			}
		}(i)
	}
	wg.Wait()
	end := time.Now()
	loggers.DefaultPrintln("\tend time: ", end)
	loggers.DefaultPrintln("\tduration: ", end.Sub(cur))
	Close()
}
