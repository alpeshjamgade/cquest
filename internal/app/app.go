package app

import (
	"context"
	"cquest/config"
	"cquest/internal/clients/db"
	"cquest/internal/constants"
	"cquest/internal/logger"
	"cquest/internal/utils"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	handler "cquest/internal/handler/cpu"
	repo "cquest/internal/repo/cpu"
	service "cquest/internal/service/cpu"
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

	Logger := logger.CreateLoggerWithCtx(ctx)

	r := GetRouter()

	db := getClients(ctx)
	cpuRepo := repo.NewCPURepo(db)
	cpuService := service.NewCPUService(cpuRepo)
	cpuHandler := handler.NewCPUHandler(cpuService)

	cpuHandler.SetupRoutes(r)

	go func() {
		Logger.Infof("Starting server on port %s", config.HttpPort)
		http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort), r)
	}()

	<-ctx.Done()
	Logger.Infof("Shutting down server...")
}

func getClients(ctx context.Context) db.DB {
	Logger := logger.CreateLoggerWithCtx(ctx)

	db := db.NewPostgresDB(config.DbHost, config.DbPort, config.DbUsername, config.DbPassword, config.DbSSLMode, config.DbName)
	err := db.Connect(ctx)

	if err != nil {
		Logger.Panic(err)
		os.Exit(1)
	}

	return db
}
