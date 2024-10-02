package handler

import (
	"context"
	"cquest/internal/constants"
	"cquest/internal/models"
	"cquest/internal/utils"
	"net/http"
	"strconv"
)

func (h *GPUHandler) GetAllGPUs(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	res := utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	gpus, err := h.service.GetAllGPUs(ctx)
	if err != nil {
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Data = gpus
	res.Status = "success"

	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *GPUHandler) AddGPU(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	req := &models.GPU{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}

	err := utils.ReadJSON(w, r, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.service.AddGPU(ctx, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *GPUHandler) UpdateGPU(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	req := &models.GPU{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	err := utils.ReadJSON(w, r, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.service.UpdateGPU(ctx, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *GPUHandler) DeleteGPUByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	queryParams := r.URL.Query()
	idStr := queryParams.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}

	err = h.service.DeleteGPUByID(ctx, id)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Message = "GPU deleted successfully"
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *GPUHandler) GetGPUByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	queryParams := r.URL.Query()
	idStr := queryParams.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	gpu, err := h.service.GetGPUByID(ctx, id)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Data = gpu
	utils.WriteJSON(w, http.StatusOK, res)
	return
}
