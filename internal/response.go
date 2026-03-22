package internal

import "bff-dashboard-api/internal/domain"

type DashboardResponse struct {
	ID               int     `json:"id"`
	FullName         string  `json:"full_name"`
	Status           string  `json:"status"`
	PendingTaskCount int     `json:"pending_task_count"`
	NextUrgentTask   *string `json:"next_urgent_task"`
	ErrorWarning     *string `json:"error_warning"`
}

func NewDashboardResponse(user domain.User, todos domain.TodoList, todosErr error) DashboardResponse {
	resp := DashboardResponse{
		ID:       user.ID,
		FullName: user.FullName(),
		Status:   user.Status(),
	}

	if todosErr != nil {
		warning := "Todos Unavailable"
		resp.ErrorWarning = &warning
		return resp
	}

	pending := todos.Pending()
	resp.PendingTaskCount = len(pending)
	resp.NextUrgentTask = todos.FirstPendingTitle()

	return resp
}
