package handler

import (
	"cquest/internal/middlewares"
	"cquest/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type CPUHandler struct {
	service service.ICPUService
}

type GPUHandler struct {
	service service.IGPUService
}

func NewCPUHandler(service service.ICPUService) *CPUHandler {
	return &CPUHandler{service: service}
}

func NewGPUHandler(service service.IGPUService) *GPUHandler {
	return &GPUHandler{service: service}
}

func (h *CPUHandler) SetupRoutes(r *mux.Router) {
	r.Use(middlewares.RequestLogger)
	r.HandleFunc("/api/cpu/all", h.GetAllCPUs).Methods(http.MethodGet)
	r.HandleFunc("/api/cpu", h.AddCPU).Methods(http.MethodPost)
	r.HandleFunc("/api/cpu", h.UpdateCPU).Methods(http.MethodPut)
	r.HandleFunc("/api/cpu", h.DeleteCPUByID).Methods(http.MethodDelete)
	r.HandleFunc("/api/cpu", h.GetCPUByID).Methods(http.MethodGet)
}

func (h *GPUHandler) SetupRoutes(r *mux.Router) {
	r.Use(middlewares.RequestLogger)
	r.HandleFunc("/api/gpu/all", h.GetAllGPUs).Methods(http.MethodGet)
	r.HandleFunc("/api/gpu", h.AddGPU).Methods(http.MethodPost)
	r.HandleFunc("/api/gpu", h.UpdateGPU).Methods(http.MethodPut)
	r.HandleFunc("/api/gpu", h.DeleteGPUByID).Methods(http.MethodDelete)
	r.HandleFunc("/api/gpu", h.GetGPUByID).Methods(http.MethodGet)

}
