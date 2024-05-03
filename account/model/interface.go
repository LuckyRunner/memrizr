package model

import (
	"context"

	"github.com/google/uuid"
)

// UserService defines methos the handler layer expects
// any service it interacts with to implement
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error)
	Signup(ctx context.Context, u *User) error
}

// TokenService defiens methods the handler layer expects to interact
// with in regards to producting JWTs as string
type TokenService interface {
	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
}

// UserRepository defines methods the service layer expects
// any repository it intects with to implement
type UserRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*User, error)
	Create(ctx context.Context, u *User) error
}
