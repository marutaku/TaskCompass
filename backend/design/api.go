package design

import (
	. "goa.design/goa/v3/dsl"
)

var Todo = ResultType("Todo", func() {
	Description("Todo describes a todo item")

	Attributes(func() {
		Attribute("id", UInt32, "Unique ID")
		Attribute("name", String, "Name of the todo")
		Attribute("description", String, "Description of the todo")
		Attribute("done", Boolean, "Whether or not the todo is done")
		Attribute("doneAt", String, "When the todo was done in ISO format")
		Attribute("doneBy", String, "Who did the todo")
		Attribute("createdAt", String, "When the todo was created in ISO format")
		Attribute("createdBy", String, "Who created the todo")
	})

	Required("name", "description", "done", "createdAt")
})

var _ = API("todo", func() {
	Title("Todo Service")
	Description("Service for managing todos")
	Server("todo", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})
