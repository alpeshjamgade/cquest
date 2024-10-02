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

	cpuHandler "cquest/internal/handler/cpu"
	cpuRepo "cquest/internal/repo/cpu"
	cpuService "cquest/internal/service/cpu"

	gpuHandler "cquest/internal/handler/gpu"
	gpuRepo "cquest/internal/repo/gpu"
	gpuService "cquest/internal/service/gpu"
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
	cpuRepo := cpuRepo.NewCPURepo(db)
	cpuService := cpuService.NewCPUService(cpuRepo)
	cpuHandler := cpuHandler.NewCPUHandler(cpuService)
	cpuHandler.SetupRoutes(r)

	gpuRepo := gpuRepo.NewGPURepo(db)
	gpuService := gpuService.NewGPUService(gpuRepo)
	gpuHandler := gpuHandler.NewGPUHandler(gpuService)
	gpuHandler.SetupRoutes(r)

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
