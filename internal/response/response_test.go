package response

import (
	"bff-dashboard-api/internal/domain"
	"errors"
	"testing"
)

func TestNewDashboardResponse_HappyPath(t *testing.T) {
	user := domain.User{ID: 1, FirstName: "Terry", LastName: "Medhurst", Age: 51}
	todos := domain.TodoList{
		{ID: 1, Title: "Done task", Completed: true},
		{ID: 2, Title: "Write tests", Completed: false},
		{ID: 3, Title: "Deploy app", Completed: false},
	}

	resp := NewDashboardResponse(user, todos, nil)

	if resp.ID != 1 {
		t.Errorf("ID: want 1, got %d", resp.ID)
	}
	if resp.FullName != "Terry Medhurst" {
		t.Errorf("FullName: want 'Terry Medhurst', got %q", resp.FullName)
	}
	if resp.Status != "Veteran" {
		t.Errorf("Status: want 'Veteran', got %q", resp.Status)
	}
	if resp.PendingTaskCount != 2 {
		t.Errorf("PendingTaskCount: want 2, got %d", resp.PendingTaskCount)
	}
	if resp.NextUrgentTask == nil {
		t.Fatal("NextUrgentTask: want a title, got nil")
	}
	if *resp.NextUrgentTask != "Write tests" {
		t.Errorf("NextUrgentTask: want 'Write tests', got %q", *resp.NextUrgentTask)
	}
	if resp.ErrorWarning != nil {
		t.Errorf("ErrorWarning: want nil, got %q", *resp.ErrorWarning)
	}
}

func TestNewDashboardResponse_TodosError(t *testing.T) {
	user := domain.User{ID: 1, FirstName: "Terry", LastName: "Medhurst", Age: 51}

	resp := NewDashboardResponse(user, nil, errors.New("timeout"))

	if resp.ErrorWarning == nil {
		t.Fatal("ErrorWarning: want a warning, got nil")
	}
	if *resp.ErrorWarning != "Todos Unavailable" {
		t.Errorf("ErrorWarning: want 'Todos Unavailable', got %q", *resp.ErrorWarning)
	}
	if resp.PendingTaskCount != 0 {
		t.Errorf("PendingTaskCount: want 0 on error, got %d", resp.PendingTaskCount)
	}
	if resp.NextUrgentTask != nil {
		t.Errorf("NextUrgentTask: want nil on error, got %q", *resp.NextUrgentTask)
	}
}

func TestNewDashboardResponse_EmptyTodos(t *testing.T) {
	user := domain.User{ID: 1, FirstName: "Terry", LastName: "Medhurst", Age: 51}

	resp := NewDashboardResponse(user, domain.TodoList{}, nil)

	if resp.PendingTaskCount != 0 {
		t.Errorf("PendingTaskCount: want 0, got %d", resp.PendingTaskCount)
	}
	if resp.NextUrgentTask != nil {
		t.Errorf("NextUrgentTask: want nil when no todos, got %q", *resp.NextUrgentTask)
	}
	if resp.ErrorWarning != nil {
		t.Errorf("ErrorWarning: want nil, got %q", *resp.ErrorWarning)
	}
}

func TestNewDashboardResponse_AllTodosCompleted(t *testing.T) {
	user := domain.User{ID: 1, FirstName: "Terry", LastName: "Medhurst", Age: 51}
	todos := domain.TodoList{
		{ID: 1, Title: "Done", Completed: true},
		{ID: 2, Title: "Also done", Completed: true},
	}

	resp := NewDashboardResponse(user, todos, nil)

	if resp.PendingTaskCount != 0 {
		t.Errorf("PendingTaskCount: want 0, got %d", resp.PendingTaskCount)
	}
	if resp.NextUrgentTask != nil {
		t.Errorf("NextUrgentTask: want nil when all completed, got %q", *resp.NextUrgentTask)
	}
}
