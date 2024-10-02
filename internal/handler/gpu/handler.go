package gpu

import (
	"cquest/internal/middlewares"
	"cquest/internal/service/gpu"
	"github.com/gorilla/mux"
	"net/http"
)

type GPUHandler struct {
	service gpu.IGPUService
}

func NewGPUHandler(service gpu.IGPUService) *GPUHandler {
	return &GPUHandler{service: service}
}

func (h *GPUHandler) SetupRoutes(r *mux.Router) {
	r.Use(middlewares.RequestLogger)
	r.HandleFunc("/api/gpu/all", h.GetAllGPUs).Methods(http.MethodGet)
	r.HandleFunc("/api/gpu", h.AddGPU).Methods(http.MethodPost)
	r.HandleFunc("/api/gpu", h.UpdateGPU).Methods(http.MethodPut)
	r.HandleFunc("/api/gpu", h.DeleteGPUByID).Methods(http.MethodDelete)
	r.HandleFunc("/api/gpu", h.GetGPUBuID).Methods(http.MethodGet)

}
