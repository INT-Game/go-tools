package gt_sys

import (
	"context"
	"github.com/INT-Game/go-tools/slog"
	"os"
	"os/signal"
	"syscall"
)

var l = slog.NewSLogger("[sys] %s")

type StopFunc func()

func StopWatcher(ctx context.Context, stop StopFunc) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP, syscall.SIGKILL)
	for {
		select {
		case <-ctx.Done():
			l.CWarn(ctx, "ctx.Done()")
			stop()
			return
		case s := <-sig:
			switch s {
			case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGKILL:
				l.CWarn(ctx, "Got %+v", s)
			case syscall.SIGHUP:
				l.CWarn(ctx, "Got syscall.SIGHUP")
			default:
				l.CWarn(ctx, "Got unknown signal")
			}
			stop()
			return
		}
	}
}
