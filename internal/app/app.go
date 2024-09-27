package app

import (
	"context"
	"cquest/config"
	"cquest/internal/constants"
	"cquest/internal/logger"
	"cquest/internal/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	err := config.LoadConf()
	if err != nil {
		fmt.Printf("Error loading config, %s", err)
		os.Exit(1)
	}

	ctx := context.WithValue(context.Background(), constants.TraceID, utils.GetUUID())
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	log := logger.CreateLoggerWithCtx(ctx)
	r := mux.NewRouter()

	go func() {
		log.Infof("Starting server on port %s", config.HttpPort)
		http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort), r)
	}()

	<-ctx.Done()
	log.Infof("Shutting down server...")
}
