package context

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/o5h/services/config"
	"github.com/o5h/services/db"
)

func Init(cfg *config.Config) (context.Context, context.CancelFunc) {
	rootContext := context.WithValue(context.Background(), config.ContextKey, cfg)
	ctx, cancel := context.WithCancel(rootContext)
	cancel = func() {
		db.Close()
		cancel()
	}
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		select {
		case s := <-c:
			log.Println("Interrupted by syscall", s)
			cancel()
			break
		case <-ctx.Done():
			log.Println("Shutdown by context done")
			break
		}
	}()
	return ctx, cancel
}
