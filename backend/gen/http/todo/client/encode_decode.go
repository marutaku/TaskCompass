// Code generated by goa v3.14.1, DO NOT EDIT.
//
// todo HTTP client encoders and decoders
//
// Command:
// $ goa gen backend/design

package client

import (
	todo "backend/gen/todo"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "todo" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListTodoPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("todo", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListRequest returns an encoder for requests sent to the todo list
// server.
func EncodeListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*todo.ListPayload)
		if !ok {
			return goahttp.ErrInvalidType("todo", "list", "*todo.ListPayload", v)
		}
		values := req.URL.Query()
		if p.Limit != nil {
			values.Add("limit", fmt.Sprintf("%v", *p.Limit))
		}
		if p.Offset != nil {
			values.Add("offset", fmt.Sprintf("%v", *p.Offset))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeListResponse returns a decoder for responses returned by the todo list
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("todo", "list", err)
			}
			err = ValidateListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("todo", "list", err)
			}
			res := NewListResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("todo", "list", resp.StatusCode, string(body))
		}
	}
}

// unmarshalTodoResponseBodyToTodoTodo builds a value of type *todo.Todo from a
// value of type *TodoResponseBody.
func unmarshalTodoResponseBodyToTodoTodo(v *TodoResponseBody) *todo.Todo {
	if v == nil {
		return nil
	}
	res := &todo.Todo{
		ID:          v.ID,
		Name:        *v.Name,
		Description: *v.Description,
		Done:        *v.Done,
		DoneAt:      v.DoneAt,
		DoneBy:      v.DoneBy,
		CreatedAt:   *v.CreatedAt,
		CreatedBy:   v.CreatedBy,
	}

	return res
}