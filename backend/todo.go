package todoapi

import (
	todo "backend/gen/todo"
	"context"
	"log"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	logger *log.Logger
}

// NewTodo returns the todo service implementation.
func NewTodo(logger *log.Logger) todo.Service {
	return &todosrvc{logger}
}

// List all todos
func (s *todosrvc) List(ctx context.Context, p *todo.ListPayload) (res *todo.ListResult, err error) {
	res = &todo.ListResult{}
	s.logger.Print("todo.list")
	return
}
