// Code generated by goa v3.14.1, DO NOT EDIT.
//
// auth service
//
// Command:
// $ goa gen backend/design

package auth

import (
	"context"
)

// The auth service manages authentication
type Service interface {
	// Login to the system
	Login(context.Context, *LoginPayload) (res *LoginResult, err error)
	// Register a new user
	Register(context.Context, *RegisterPayload) (res *RegisterResult, err error)
	// Logout of the system
	Logout(context.Context, *LogoutPayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "auth"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"login", "register", "logout"}

// LoginPayload is the payload type of the auth service login method.
type LoginPayload struct {
	// Username to login with
	Username string
	// Password to login with
	Password string
}

// LoginResult is the result type of the auth service login method.
type LoginResult struct {
	// JWT token to use for authentication
	Token *string
}

// LogoutPayload is the payload type of the auth service logout method.
type LogoutPayload struct {
	// JWT token to use for authentication
	Token string
}

// RegisterPayload is the payload type of the auth service register method.
type RegisterPayload struct {
	// Username to register with
	Username string
	// Password to register with
	Password string
}

// RegisterResult is the result type of the auth service register method.
type RegisterResult struct {
	// JWT token to use for authentication
	Token *string
}
