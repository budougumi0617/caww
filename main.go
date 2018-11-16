// Copyright © 2018 budougumi0617 All Rights Reserved.

package main

import (
	"log"
	"net/http"
)

func main() {
	router := GetTodoRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
