// +build e2e

// Copyright © 2019 budougumi0617 All Rights Reserved.
package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/budougumi0617/caww/repository"
	"github.com/budougumi0617/caww/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func TestUserCase_Save(t *testing.T) {
	okName := "budougumi0617"
	okEmail := "budougumi0617@example.com"
	type args struct {
		name, email string
	}
	okArgs := args{
		name:  okName,
		email: okEmail,
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Correct",
			args: okArgs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := os.Getenv("MYSQL_USER")
			p := os.Getenv("MYSQL_PORT")
			db, err := sql.Open("mysql", fmt.Sprintf("%s:@(localhost:%s)/caww?parseTime=true&loc=Asia%%2FTokyo", u, p))
			if err != nil {
				log.Fatalln(err)
			}
			ctx := context.Background()
			repo := repository.NewRepo(db)
			uc := usecase.NewUserCase(repo)

			got, err := uc.Save(ctx, tt.args.name, tt.args.email)

			if err != nil {
				t.Errorf("want no err, but has error %#v", err)
			}

			if got == 0 {
				t.Error("ID was 0")
			}
		})
	}
}
