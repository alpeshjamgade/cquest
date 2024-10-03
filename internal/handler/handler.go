package handler

import (
	"cquest/internal/middlewares"
	"cquest/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	Service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) SetupRoutes(r *mux.Router) {
	r.Use(middlewares.RequestLogger)

	// cpu
	r.HandleFunc("/api/cpu/all", h.GetAllCPUs).Methods(http.MethodGet)
	r.HandleFunc("/api/cpu", h.AddCPU).Methods(http.MethodPost)
	r.HandleFunc("/api/cpu", h.UpdateCPU).Methods(http.MethodPut)
	r.HandleFunc("/api/cpu", h.DeleteCPUById).Methods(http.MethodDelete)
	r.HandleFunc("/api/cpu", h.GetCPUById).Methods(http.MethodGet)

	// gpu
	r.HandleFunc("/api/gpu/all", h.GetAllGPUs).Methods(http.MethodGet)
	r.HandleFunc("/api/gpu", h.AddGPU).Methods(http.MethodPost)
	r.HandleFunc("/api/gpu", h.UpdateGPU).Methods(http.MethodPut)
	r.HandleFunc("/api/gpu", h.DeleteGPUById).Methods(http.MethodDelete)
	r.HandleFunc("/api/gpu", h.GetGPUById).Methods(http.MethodGet)

	// laptop
	r.HandleFunc("/api/laptop/all", h.GetAllLaptops).Methods(http.MethodGet)
	r.HandleFunc("/api/laptop", h.AddLaptop).Methods(http.MethodPost)
	r.HandleFunc("/api/laptop", h.UpdateLaptop).Methods(http.MethodPut)
	r.HandleFunc("/api/laptop", h.DeleteLaptopById).Methods(http.MethodDelete)
	r.HandleFunc("/api/laptop", h.GetLaptopById).Methods(http.MethodGet)
	r.HandleFunc("/api/laptop/cpu/{cpuId}", h.GetLaptopsByCPUId).Methods(http.MethodGet)
	r.HandleFunc("/api/laptop/gpu/{gpuId}", h.GetLaptopsByGPUId).Methods(http.MethodGet)
}
