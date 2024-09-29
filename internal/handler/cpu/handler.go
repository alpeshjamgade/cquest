package cpu

import (
	"cquest/internal/middlewares"
	"cquest/internal/service/cpu"
	"github.com/gorilla/mux"
	"net/http"
)

type CPUHandler struct {
	service cpu.ICPUService
}

func NewCPUHandler(service cpu.ICPUService) *CPUHandler {
	return &CPUHandler{service: service}
}

func (h *CPUHandler) SetupRoutes(r *mux.Router) {
	r.Use(middlewares.RequestLogger)
	r.HandleFunc("/api/cpu/all", h.GetAllCPUs).Methods(http.MethodGet)
	r.HandleFunc("/api/cpu", h.AddCPU).Methods(http.MethodPost)
	r.HandleFunc("/api/cpu", h.UpdateCPU).Methods(http.MethodPut)
	r.HandleFunc("/api/cpu", h.DeleteCPUByID).Methods(http.MethodDelete)
	r.HandleFunc("/api/cpu", h.GetCPUByID).Methods(http.MethodGet)
}
