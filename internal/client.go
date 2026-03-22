package internal

import (
	"bff-dashboard-api/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func NewDummyJSONClient(httpClient *http.Client, baseURL string) *DummyJSONClient {
	return &DummyJSONClient{
		httpClient: httpClient,
		baseURL:    strings.TrimRight(baseURL, "/"),
	}
}

func (c *DummyJSONClient) GetUser(ctx context.Context, id int) (domain.User, error) {
	var env userEnvelope
	if err := c.getJSON(ctx, fmt.Sprintf("%s/users/%d", c.baseURL, id), &env); err != nil {
		return domain.User{}, fmt.Errorf("GetUser %d: %w", id, err)
	}
	return domain.User{
		ID:        env.ID,
		FirstName: env.FirstName,
		LastName:  env.LastName,
		Age:       env.Age,
	}, nil
}

func (c *DummyJSONClient) GetTodos(ctx context.Context, id int) (domain.TodoList, error) {
	var env todosEnvelope
	if err := c.getJSON(ctx, fmt.Sprintf("%s/todos/user/%d?delay=3000", c.baseURL, id), &env); err != nil {
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
