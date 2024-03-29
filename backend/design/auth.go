package design

import . "goa.design/goa/v3/dsl"

var _ = Service("auth", func() {
	Description("The auth service manages authentication")
	Method("login", func() {
		Description("Login to the system")
		Payload(func() {
			Attribute("username", String, "Username to login with")
			Attribute("password", String, "Password to login with")
			Required("username", "password")
		})
		Result(func() {
			Attribute("token", String, "JWT token to use for authentication")
		})
		HTTP(func() {
			POST("/login")
			Response(StatusOK)
		})
	})
	Method("register", func() {
		Description("Register a new user")
		Payload(func() {
			Attribute("username", String, "Username to register with")
			Attribute("password", String, "Password to register with")
			Required("username", "password")
		})
		Result(func() {
			Attribute("token", String, "JWT token to use for authentication")
		})
		HTTP(func() {
			POST("/register")
			Response(StatusOK)
		})
	})
	Method("logout", func() {
		Description("Logout of the system")
		Payload(func() {
			Attribute("token", String, "JWT token to use for authentication")
			Required("token")
		})
		HTTP(func() {
			POST("/logout")
			Response(StatusOK)
		})
	})
})
