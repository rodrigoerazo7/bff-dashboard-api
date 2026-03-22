# Dashboard API BFF

## Checklist

- [x] Base server with GET `dashboard/:id` endpoint.
- [ ] Client, external data source: `/users/:id`
- [ ] Client, external data source: `/todos/user/:id`
- [ ] Client, strict 2 seconds timeout.
- [ ] Test client `user`
- [ ] Test client `todo`
- [ ] **Fetch calls in parallel.**
- [x] Business rule, full name (first + last).
- [x] Business rule, age > 50 with status Veteran, otherwise Rookie.
- [ ] Business rule, count of pending tasks.
- [x] Business rule, title of the first pending task.
- [x] Business rule, filter todos completed.
- [ ] Business rule, partial failure handling with warning.
- [x] Test and validate business rules (QA).
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
