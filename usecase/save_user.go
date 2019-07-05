// Copyright © 2019 budougumi0617 All Rights Reserved.

package usecase

import (
	"context"

	"github.com/google/wire"

	"github.com/budougumi0617/caww/entity"
	"github.com/budougumi0617/caww/usecase/port"
)

// UserSearcherSet returns filled UserCase object.
var UserSearcherSet = wire.NewSet(
	NewUserSearcher,
)

// UserSearcher :
// TODO SaveとFindで分ける
type UserSearcher struct {
	ua port.UserReader
}

// NewUserSearcher :
func NewUserSearcher(ua port.UserReader) *UserSearcher {
	return &UserSearcher{
		ua: ua,
	}
}

// Find :
func (au *UserSearcher) Find(ctx context.Context, id int64) (*entity.User, error) {
	u, err := au.ua.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
