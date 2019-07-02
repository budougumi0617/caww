// Copyright © 2019 budougumi0617 All Rights Reserved.

package usecase

import (
	"context"

	"github.com/google/wire"

	"github.com/budougumi0617/caww/entity"
	"github.com/budougumi0617/caww/usecase/port"
)

// UserCaseSet returns filled UserCase object.
var UserCaseSet = wire.NewSet(
	NewUserCase,
)

// UserCase :
// TODO SaveとFindで分ける
type UserCase struct {
	ua port.UserAccessor
}

// NewUserCase :
func NewUserCase(ua port.UserAccessor) *UserCase {
	return &UserCase{
		ua: ua,
	}
}

// Save :
func (au *UserCase) Save(ctx context.Context, name, email string) (int64, error) {
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

// Find :
func (au *UserCase) Find(ctx context.Context, id int64) (*entity.User, error) {
	u, err := au.ua.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
