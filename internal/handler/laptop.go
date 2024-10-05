package handler

import (
	"context"
	"cquest/internal/constants"
	"cquest/internal/models"
	"cquest/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllLaptops(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	laptops, err := h.Service.GetAllLaptops(ctx)
	if err != nil {
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusInternalServerError, res)
		return
	}

	res.Data = laptops
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) AddLaptop(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	req := &models.Laptop{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Status: "error", Message: ""}
	err := utils.ReadJSON(w, r, req)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.Service.AddLaptop(ctx, req)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Status = "success"
	res.Message = "Laptop added successfully"

	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *Handler) UpdateLaptop(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	req := &models.Laptop{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Status: "error", Message: ""}
	err := utils.ReadJSON(w, r, req)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateLaptop(ctx, req)
	if err != nil {
		res.Message = err.Error()
		res.Status = "error"
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Status = "success"
	res.Message = "Laptop updated successfully"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) DeleteLaptopById(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	queryParams := r.URL.Query()
	res := &utils.HTTPResponse{Data: map[string]string{}, Status: "error", Message: ""}
	idStr := queryParams.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteLaptopById(ctx, id)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Message = "Laptop deleted successfully"
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) GetLaptopById(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	queryParams := r.URL.Query()
	res := &utils.HTTPResponse{Data: map[string]string{}, Status: "error", Message: ""}
	idStr := queryParams.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	laptop, err := h.Service.GetLaptopById(ctx, id)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Data = laptop
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) GetLaptopsByCPUId(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	pathParams := mux.Vars(r)
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	cpuIdStr := pathParams["cpu_id"]
	cpuId, err := strconv.Atoi(cpuIdStr)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	laptop, err := h.Service.GetLaptopsByCPUId(ctx, cpuId)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
	}

	res.Status = "success"
	res.Data = laptop
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) GetLaptopsByGPUId(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	pathParams := mux.Vars(r)
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	gpuIdStr := pathParams["gpu_id"]
	gpuId, err := strconv.Atoi(gpuIdStr)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	laptop, err := h.Service.GetLaptopsByCPUId(ctx, gpuId)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
	}

	res.Status = "success"
	res.Data = laptop
	utils.WriteJSON(w, http.StatusOK, res)
	return
}
