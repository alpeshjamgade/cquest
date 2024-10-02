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
	r.HandleFunc("/api/cpu", h.DeleteCPUByID).Methods(http.MethodDelete)
	r.HandleFunc("/api/cpu", h.GetCPUByID).Methods(http.MethodGet)

	// gpu
	r.HandleFunc("/api/gpu/all", h.GetAllGPUs).Methods(http.MethodGet)
	r.HandleFunc("/api/gpu", h.AddGPU).Methods(http.MethodPost)
	r.HandleFunc("/api/gpu", h.UpdateGPU).Methods(http.MethodPut)
	r.HandleFunc("/api/gpu", h.DeleteGPUByID).Methods(http.MethodDelete)
	r.HandleFunc("/api/gpu", h.GetGPUByID).Methods(http.MethodGet)
}
