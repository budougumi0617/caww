// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/budougumi0617/caww/repository"
	"github.com/budougumi0617/caww/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@(localhost:43306)/sqlx_development?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	repo := repository.NewRepo(db)
	u, err := repo.FindUser(ctx, 1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("u = %+v\n", u)
	name := "budougumi0617"
	email := "budougumi0617@example.com"
	au := usecase.NewUserCase(repo)
	id, err := au.Save(ctx, name, email)
	if err != nil {
		log.Fatalln(err)
	}
	u, err = au.Find(ctx, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("readed u = %+v\n", u)

	// router := GetTodoRouter()
	// log.Fatal(http.ListenAndServe(":8080", router))
}
