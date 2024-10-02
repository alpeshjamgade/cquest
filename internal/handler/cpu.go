package handler

import (
	"context"
	"cquest/internal/constants"
	"cquest/internal/logger"
	"cquest/internal/models"
	"cquest/internal/utils"
	"net/http"
	"strconv"
)

func (h *CPUHandler) GetAllCPUs(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateLoggerWithCtx(ctx)
	result, err := h.service.GetAllCPUs(ctx)
	if err != nil {
		Logger.Errorw("error while getting all cpus", "error", err)
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res := utils.HTTPResponse{Data: result, Status: "success", Message: ""}
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *CPUHandler) AddCPU(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateLoggerWithCtx(ctx)

	req := &models.CPU{}
	res := utils.HTTPResponse{Data: map[string]string{}, Status: "success", Message: ""}

	err := utils.ReadJSON(w, r, req)
	if err != nil {
		Logger.Errorw("error reading request", "err", err)
		res.Status = "error"
		res.Message = "Invalid Request"
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	err = h.service.AddCPU(ctx, req)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	res.Status = "success"
	res.Message = "successfully added CPU"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *CPUHandler) UpdateCPU(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateLoggerWithCtx(ctx)
	req := &models.CPU{}
	res := utils.HTTPResponse{Data: map[string]string{}, Status: "success", Message: ""}
	err := utils.ReadJSON(w, r, req)
	if err != nil {
		Logger.Errorw("error reading request", "err", err)
		res.Status = "error"
		res.Message = "Invalid Request"
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}
	err = h.service.UpdateCPU(ctx, req)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}
	res.Status = "success"
	res.Message = "successfully updated CPU"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *CPUHandler) DeleteCPUByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	res := utils.HTTPResponse{Data: map[string]string{}, Status: "success", Message: ""}
	queryParams := r.URL.Query()
	idStr := queryParams.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	err = h.service.DeleteCPUByID(ctx, id)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}
	res.Status = "success"
	res.Message = "successfully deleted CPU"
	utils.WriteJSON(w, http.StatusOK, res)

	return
}

func (h *CPUHandler) GetCPUByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	res := utils.HTTPResponse{Data: map[string]string{}, Status: "success", Message: ""}

	queryParams := r.URL.Query()
	idStr := queryParams.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	cpu, err := h.service.GetCPUByID(ctx, id)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}
	res.Data = cpu
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}
