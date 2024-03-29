// Code generated by goa v3.14.1, DO NOT EDIT.
//
// auth HTTP server types
//
// Command:
// $ goa gen backend/design

package server

import (
	auth "backend/gen/auth"

	goa "goa.design/goa/v3/pkg"
)

// LoginRequestBody is the type of the "auth" service "login" endpoint HTTP
// request body.
type LoginRequestBody struct {
	// Username to login with
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	// Password to login with
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// RegisterRequestBody is the type of the "auth" service "register" endpoint
// HTTP request body.
type RegisterRequestBody struct {
	// Username to register with
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	// Password to register with
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// LogoutRequestBody is the type of the "auth" service "logout" endpoint HTTP
// request body.
type LogoutRequestBody struct {
	// JWT token to use for authentication
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// LoginResponseBody is the type of the "auth" service "login" endpoint HTTP
// response body.
type LoginResponseBody struct {
	// JWT token to use for authentication
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// RegisterResponseBody is the type of the "auth" service "register" endpoint
// HTTP response body.
type RegisterResponseBody struct {
	// JWT token to use for authentication
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
}

// NewLoginResponseBody builds the HTTP response body from the result of the
// "login" endpoint of the "auth" service.
func NewLoginResponseBody(res *auth.LoginResult) *LoginResponseBody {
	body := &LoginResponseBody{
		Token: res.Token,
	}
	return body
}

// NewRegisterResponseBody builds the HTTP response body from the result of the
// "register" endpoint of the "auth" service.
func NewRegisterResponseBody(res *auth.RegisterResult) *RegisterResponseBody {
	body := &RegisterResponseBody{
		Token: res.Token,
	}
	return body
}

// NewLoginPayload builds a auth service login endpoint payload.
func NewLoginPayload(body *LoginRequestBody) *auth.LoginPayload {
	v := &auth.LoginPayload{
		Username: *body.Username,
		Password: *body.Password,
	}

	return v
}

// NewRegisterPayload builds a auth service register endpoint payload.
func NewRegisterPayload(body *RegisterRequestBody) *auth.RegisterPayload {
	v := &auth.RegisterPayload{
		Username: *body.Username,
		Password: *body.Password,
	}

	return v
}

// NewLogoutPayload builds a auth service logout endpoint payload.
func NewLogoutPayload(body *LogoutRequestBody) *auth.LogoutPayload {
	v := &auth.LogoutPayload{
		Token: *body.Token,
	}

	return v
}

// ValidateLoginRequestBody runs the validations defined on LoginRequestBody
func ValidateLoginRequestBody(body *LoginRequestBody) (err error) {
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}

// ValidateRegisterRequestBody runs the validations defined on
// RegisterRequestBody
func ValidateRegisterRequestBody(body *RegisterRequestBody) (err error) {
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.Password == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("password", "body"))
	}
	return
}

// ValidateLogoutRequestBody runs the validations defined on LogoutRequestBody
func ValidateLogoutRequestBody(body *LogoutRequestBody) (err error) {
	if body.Token == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("token", "body"))
	}
	return
}
