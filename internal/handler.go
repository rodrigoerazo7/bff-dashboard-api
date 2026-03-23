package internal

import (
	"bff-dashboard-api/internal/domain"
	"bff-dashboard-api/internal/response"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type DashboardServiceInterface interface {
	BuildDashboard(ctx context.Context, id int) (domain.User, domain.TodoList, error)
}

type DashboardHTTPHandler struct {
	service DashboardServiceInterface
}

func NewDashboardHandler(service DashboardServiceInterface) *DashboardHTTPHandler {
	return &DashboardHTTPHandler{service: service}
}

func (h *DashboardHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/dashboard/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, todos, err := h.service.BuildDashboard(r.Context(), id)
	if user.ID == 0 {
		slog.Error("failed to fetch user", "id", id, "err", err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	resp := response.NewDashboardResponse(user, todos, err)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
