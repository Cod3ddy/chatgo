package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Cod3ddy/chatgo/chat/foundation/logger"
)

func main() {
	var log *logger.Logger

	traceIDFn := func(ctx context.Context) string{
		return "" // NEED TRACE 
	}


	log = logger.New(os.Stdout, slog.Level(logger.LevelDebug), "CAP", traceIDFn)

	ctx := context.Background()

	if err := run(ctx, log); err != nil{
		log.Error(ctx, "startup", "error",err)
		os.Exit(1)
	}
}


func run(ctx context.Context, log *logger.Logger) error{
	// ------------------------------------------------------------
	//GOMAXPROCS

	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))


	// ------------------------------------------
	
	log.Info(ctx, "status", "started")

	defer 	log.Info(ctx, "status", "shutting down")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	return nil
}