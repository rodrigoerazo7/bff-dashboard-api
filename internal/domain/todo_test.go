package domain

import "testing"

func TestTodoList_Pending_Mixed(t *testing.T) {
	todos := TodoList{
		{ID: 1, Completed: true},
		{ID: 2, Completed: false},
		{ID: 3, Completed: false},
	}
	if got := len(todos.Pending()); got != 2 {
		t.Errorf("Pending(): want 2, got %d", got)
	}
}

func TestTodoList_Pending_Empty(t *testing.T) {
	if got := len(TodoList{}.Pending()); got != 0 {
		t.Errorf("Pending(): want 0 on empty list, got %d", got)
	}
}

func TestTodoList_Pending_AllCompleted(t *testing.T) {
	todos := TodoList{
		{ID: 1, Completed: true},
		{ID: 2, Completed: true},
	}
	if got := len(todos.Pending()); got != 0 {
		t.Errorf("Pending(): want 0 when all done, got %d", got)
	}
}

func TestTodoList_FirstPendingTitle_ReturnsTitleOfFirstPending(t *testing.T) {
	todos := TodoList{
		{ID: 1, Title: "Done task", Completed: true},
		{ID: 2, Title: "Do something important", Completed: false},
		{ID: 3, Title: "Another task", Completed: false},
	}
	got := todos.FirstPendingTitle()
	if got == nil {
		t.Fatal("FirstPendingTitle(): want a title, got nil")
	}
	if *got != "Do something important" {
		t.Errorf("FirstPendingTitle(): want 'Do something important', got %q", *got)
	}
}

func TestTodoList_FirstPendingTitle_NilWhenNoPending(t *testing.T) {
	todos := TodoList{
		{ID: 1, Completed: true},
	}
	if got := todos.FirstPendingTitle(); got != nil {
		t.Errorf("FirstPendingTitle(): want nil, got %q", *got)
	}
}
