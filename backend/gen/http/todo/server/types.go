// Code generated by goa v3.14.1, DO NOT EDIT.
//
// todo HTTP server types
//
// Command:
// $ goa gen backend/design

package server

import (
	todo "backend/gen/todo"
)

// ListResponseBody is the type of the "todo" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// List of todos
	Todos TodoCollectionResponseBody `form:"todos,omitempty" json:"todos,omitempty" xml:"todos,omitempty"`
}

// ShowResponseBody is the type of the "todo" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// Todo to show
	Todo *TodoResponseBody `form:"todo,omitempty" json:"todo,omitempty" xml:"todo,omitempty"`
}

// TodoCollectionResponseBody is used to define fields on response body types.
type TodoCollectionResponseBody []*TodoResponseBody

// TodoResponseBody is used to define fields on response body types.
type TodoResponseBody struct {
	// Unique ID
	ID *uint32 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Name of the todo
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the todo
	Description string `form:"description" json:"description" xml:"description"`
	// Whether or not the todo is done
	Done bool `form:"done" json:"done" xml:"done"`
	// When the todo was done in ISO format
	DoneAt *string `form:"doneAt,omitempty" json:"doneAt,omitempty" xml:"doneAt,omitempty"`
	// Who did the todo
	DoneBy *string `form:"doneBy,omitempty" json:"doneBy,omitempty" xml:"doneBy,omitempty"`
	// When the todo was created in ISO format
	CreatedAt string `form:"createdAt" json:"createdAt" xml:"createdAt"`
	// Who created the todo
	CreatedBy *string `form:"createdBy,omitempty" json:"createdBy,omitempty" xml:"createdBy,omitempty"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "todo" service.
func NewListResponseBody(res *todo.ListResult) *ListResponseBody {
	body := &ListResponseBody{}
	if res.Todos != nil {
		body.Todos = make([]*TodoResponseBody, len(res.Todos))
		for i, val := range res.Todos {
			body.Todos[i] = marshalTodoTodoToTodoResponseBody(val)
		}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "todo" service.
func NewShowResponseBody(res *todo.ShowResult) *ShowResponseBody {
	body := &ShowResponseBody{}
	if res.Todo != nil {
		body.Todo = marshalTodoTodoToTodoResponseBody(res.Todo)
	}
	return body
}

// NewListPayload builds a todo service list endpoint payload.
func NewListPayload(limit *uint32, offset *uint32) *todo.ListPayload {
	v := &todo.ListPayload{}
	v.Limit = limit
	v.Offset = offset

	return v
}

// NewShowPayload builds a todo service show endpoint payload.
func NewShowPayload(id uint32) *todo.ShowPayload {
	v := &todo.ShowPayload{}
	v.ID = id

	return v
}
