package gpu

import (
	"context"
	"cquest/internal/constants"
	"cquest/internal/models"
	"cquest/internal/utils"
	"net/http"
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
	req := models.GPU{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}

	err := utils.ReadJSON(w, r, &req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	h.service.AddGPU(ctx, req)
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *GPUHandler) UpdateGPU(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	req := models.GPU{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	err := utils.ReadJSON(w, r, &req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.service.UpdateGPU(ctx, req)
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *GPUHandler) DeleteGPUByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	req := models.GPU{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}
	err := utils.ReadJSON(w, r, &req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = h.service.DeleteGPUByID(ctx, req.ID)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Message = "GPU deleted successfully"
	res.Status = "success"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *GPUHandler) GetGPUBuID(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	req := models.GPU{}
	res := &utils.HTTPResponse{Data: map[string]string{}, Message: "", Status: "error"}

	err := utils.ReadJSON(w, r, &req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	gpu, err := h.service.GetGPUByID(ctx, req.ID)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	res.Data = gpu
	utils.WriteJSON(w, http.StatusOK, res)
	return
}
