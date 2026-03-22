package internal

import (
	"bff-dashboard-api/internal/domain"
	"context"
	"fmt"
	"sync"
)

type DashboardService struct {
	client DashboardDataClient
}

func NewDashboardService(c DashboardDataClient) *DashboardService {
	return &DashboardService{client: c}
}

func (s *DashboardService) BuildDashboard(ctx context.Context, id int) (domain.User, domain.TodoList, error) {
	var (
		wg                sync.WaitGroup
		user              domain.User
		todos             domain.TodoList
		userErr, todosErr error
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		user, userErr = s.client.GetUser(ctx, id)
	}()

	go func() {
		defer wg.Done()
		todos, todosErr = s.client.GetTodos(ctx, id)
	}()

	wg.Wait()

	if userErr != nil {
		return domain.User{}, nil, fmt.Errorf("fetch user: %w", userErr)
	}

	return user, todos, todosErr
}
