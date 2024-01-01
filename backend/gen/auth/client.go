// Code generated by goa v3.14.1, DO NOT EDIT.
//
// auth client
//
// Command:
// $ goa gen backend/design

package auth

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "auth" service client.
type Client struct {
	LoginEndpoint    goa.Endpoint
	RegisterEndpoint goa.Endpoint
	LogoutEndpoint   goa.Endpoint
}

// NewClient initializes a "auth" service client given the endpoints.
func NewClient(login, register, logout goa.Endpoint) *Client {
	return &Client{
		LoginEndpoint:    login,
		RegisterEndpoint: register,
		LogoutEndpoint:   logout,
	}
}

// Login calls the "login" endpoint of the "auth" service.
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res *LoginResult, err error) {
	var ires any
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginResult), nil
}

// Register calls the "register" endpoint of the "auth" service.
func (c *Client) Register(ctx context.Context, p *RegisterPayload) (res *RegisterResult, err error) {
	var ires any
	ires, err = c.RegisterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*RegisterResult), nil
}

// Logout calls the "logout" endpoint of the "auth" service.
func (c *Client) Logout(ctx context.Context, p *LogoutPayload) (err error) {
	_, err = c.LogoutEndpoint(ctx, p)
	return
}