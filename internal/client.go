package internal

import (
	"bff-dashboard-api/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type DashboardDataClient interface {
	GetUser(ctx context.Context, id int) (domain.User, error)
	GetTodos(ctx context.Context, id int) (domain.TodoList, error)
}

type todosEnvelope struct {
	Todos []struct {
		ID        int    `json:"id"`
		Todo      string `json:"todo"`
		Completed bool   `json:"completed"`
		UserID    int    `json:"userId"`
	} `json:"todos"`
}

type userEnvelope struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

type DummyJSONClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewDummyJSONClient(baseURL string, timeout time.Duration) *DummyJSONClient {
	return &DummyJSONClient{
		httpClient: &http.Client{Timeout: timeout},
		baseURL:    strings.TrimRight(baseURL, "/"),
	}
}

func (c *DummyJSONClient) GetUser(ctx context.Context, id int) (domain.User, error) {
	var env userEnvelope
	url := fmt.Sprintf("%s/users/%d", c.baseURL, id)
	if delay, second := simulatedDelay("GetUser"); delay > 0 {
		url += fmt.Sprintf("?delay=%d", delay)
		slog.Info("GetUser: simulating delay", "delay_ms", delay, "second", second)
	} else {
		slog.Info("GetUser: no delay", "second", second)
	}

	if err := c.getJSON(ctx, url, &env); err != nil {
		return domain.User{}, fmt.Errorf("GetUser %d: %w", id, err)
	}

	user := domain.User{
		ID:        env.ID,
		FirstName: env.FirstName,
		LastName:  env.LastName,
		Age:       env.Age,
	}
	slog.Info("GetUser result", "user", user)
	return user, nil
}

func (c *DummyJSONClient) GetTodos(ctx context.Context, id int) (domain.TodoList, error) {
	var env todosEnvelope
	url := fmt.Sprintf("%s/todos/user/%d", c.baseURL, id)
	if delay, second := simulatedDelay("GetTodos"); delay > 0 {
		url += fmt.Sprintf("?delay=%d", delay)
		slog.Info("GetTodos: simulating delay", "delay_ms", delay, "second", second)
	} else {
		slog.Info("GetTodos: no delay", "second", second)
	}

	if err := c.getJSON(ctx, url, &env); err != nil {
		return nil, fmt.Errorf("GetTodos %d: %w", id, err)
	}

	todos := make(domain.TodoList, 0, len(env.Todos))
	for _, t := range env.Todos {
		todos = append(todos, domain.Todo{
			ID:        t.ID,
			Title:     t.Todo,
			Completed: t.Completed,
			UserID:    t.UserID,
		})
	}
	slog.Info("GetTodos result", "count", len(todos), "todos", todos)
	return todos, nil
}

func (c *DummyJSONClient) getJSON(ctx context.Context, url string, target any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}

	return nil
}

func simulatedDelay(caller string) (int, int) {
	second := time.Now().Second()
	switch caller {
	case "GetUser":
		if second > 30 && second%3 == 0 {
			return 3000, second
		}
	case "GetTodos":
		if second < 30 && second%3 == 0 {
			return 3000, second
		}
	}
	return 0, second
}
