# Dashboard API BFF

## To-do list

- [ ] Base server with GET `dashboard/:id` endpoint.
- [ ] Client, external data source: `/users/:id`
- [ ] Client, external data source: `/todos/user/:id`
- [ ] Client, strict 2 seconds timeout.
- [ ] Test client `user`
- [ ] Test client `todo`
- [ ] **Fetch calls in parallel.**
- [ ] Business logic, full name (first + last).
- [ ] Business logic, age > 50 with status Veteran, otherwise Rookie.
- [ ] Business logic, count of pending tasks.
- [ ] Business logic, title of the first pending task.
- [ ] Business logic, filter todos completed.
- [ ] Business logic, partial failure handling with warning.
- [ ] Unit tests.
- [ ] README.
- [ ] GIT.

### Expected output

```json
{
  "id": 1,
  "full_name": "Terry Medhurst",
  "status": "Rookie",
  "pending_task_count": 2,
  "next_urgent_task": "Do something important",
  "error_warning": null
}
```

## External Data Sources

1. User Profile: https://dummyjson.com/users/:id
2. User Todos: https://dummyjson.com/todos/user/:id
3. User ids: 1, 2, 3, 4, 5, 6, 7, 8.
