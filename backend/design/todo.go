package design

import . "goa.design/goa/v3/dsl"

var _ = Service("todo", func() {
	Description("The todo service manages todo lists")
	Method("list", func() {
		Description("List all todos")
		Payload(func() {
			Attribute("limit", UInt32, "Maximum number of todos to return")
			Attribute("offset", UInt32, "Offset into the list of todos to start at")
		})
		Result(func() {
			Attribute("todos", CollectionOf(Todo), "List of todos")

		})
		HTTP(func() {
			GET("/todos")
			Param("limit")
			Param("offset")
			Response(StatusOK)
		})
	})
})
