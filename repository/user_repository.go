// Copyright © 2019 budougumi0617 All Rights Reserved.

package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/budougumi0617/caww/entity"
)

// FindUser implements port.UserRepository.
func (repo *Repo) FindUser(ctx context.Context, id int64) (*entity.User, error) {
	u := &entity.User{}
	err := repo.db.QueryRowContext(ctx, "SELECT id, name, email, created_at, updated_at FROM user WHERE id = ?", id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return u, nil
}

// AddUser :
func (repo *Repo) AddUser(ctx context.Context, u *entity.User) error {
	stmt, err := repo.db.PrepareContext(ctx, `
        INSERT INTO user (name, email, created_at, updated_at)
        VALUES (?, ?, ?, ?)
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()
	now := time.Now()

	res, err := stmt.ExecContext(ctx, u.Name, u.Email, now, now)
	id, err := res.LastInsertId() // 挿入した行のIDを返却
	if err != nil {
		return err
	}
	u.ID = id
	u.CreatedAt = now
	u.UpdatedAt = now

	return nil
}
