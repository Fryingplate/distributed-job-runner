package job

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *Service
	Queue   interface {
		Publish(string) error
	}
}

type CreateJobRequest struct {
	Name string `json:"name"`
}

func (h *Handler) CreateJob(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req CreateJobRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	j, err := h.Service.Create(
		req.Name,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	err = h.Queue.Publish(j.ID)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(j)
}
