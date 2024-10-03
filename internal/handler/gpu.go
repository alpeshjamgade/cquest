package handler

import (
	"context"
	"cquest/internal/constants"
	"cquest/internal/models"
	"cquest/internal/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllGPUs(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	res := utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	gpus, err := h.Service.GetAllGPUs(ctx)
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

func (h *Handler) AddGPU(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	req := &models.GPU{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}

	err := utils.ReadJSON(w, r, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.Service.AddGPU(ctx, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *Handler) UpdateGPU(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	req := &models.GPU{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	err := utils.ReadJSON(w, r, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateGPU(ctx, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) DeleteGPUById(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	queryParams := r.URL.Query()
	idStr := queryParams.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}

	err = h.Service.DeleteGPUById(ctx, id)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Message = "GPU deleted successfully"
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) GetGPUById(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	queryParams := r.URL.Query()
	idStr := queryParams.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	gpu, err := h.Service.GetGPUById(ctx, id)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Data = gpu
	utils.WriteJSON(w, http.StatusOK, res)
	return
}
