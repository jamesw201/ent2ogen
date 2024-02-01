// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddKeyboard implements addKeyboard operation.
	//
	// POST /keyboard
	AddKeyboard(ctx context.Context, req *Keyboard) (*Keyboard, error)
	// GetKeyboard implements getKeyboard operation.
	//
	// GET /keyboard/{id}
	GetKeyboard(ctx context.Context, params GetKeyboardParams) (*Keyboard, error)
	// KeyboardGet implements GET /keyboard operation.
	//
	// GET /keyboard
	KeyboardGet(ctx context.Context) ([]Keyboard, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
