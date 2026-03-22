package domain

type Todo struct {
	ID        int
	Title     string
	Completed bool
	UserID    int
}

type TodoList []Todo

// Pending returns only the todos that have not been completed.
func (tl TodoList) Pending() TodoList {
	result := make(TodoList, 0, len(tl))
	for _, t := range tl {
		if t.IsPending() {
			result = append(result, t)
		}
	}
	return result
}

// FirstPendingTitle returns the title of the first pending todo,
// or nil if there are no pending todos.
func (tl TodoList) FirstPendingTitle() *string {
	pending := tl.Pending()
	if len(pending) == 0 {
		return nil
	}
	title := pending[0].Title
	return &title
}

// IsPending returns true if the task has not been completed.
func (t Todo) IsPending() bool {
	return !t.Completed
}
