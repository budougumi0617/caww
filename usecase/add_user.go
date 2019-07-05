// Copyright © 2019 budougumi0617 All Rights Reserved.

package usecase

import (
	"context"

	"github.com/google/wire"

	"github.com/budougumi0617/caww/entity"
	"github.com/budougumi0617/caww/usecase/port"
)

// UserCreatorSet returns filled UserCase object.
var UserCreatorSet = wire.NewSet(
	NewUserCreator,
)

// UserCreator :
// TODO SaveとFindで分ける
type UserCreator struct {
	ua port.UserWriter
}

// NewUserCreator :
func NewUserCreator(ua port.UserWriter) *UserCreator {
	return &UserCreator{
		ua: ua,
	}
}

// Save :
func (au *UserCreator) Save(ctx context.Context, name, email string) (int64, error) {
	u := &entity.User{
		Name:  name,
		Email: email,
	}
	err := au.ua.AddUser(ctx, u)
	if err != nil {
		return 0, err
	}

	return u.ID, nil
}
