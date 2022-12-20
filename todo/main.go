package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tigrisdata/todo/model"
	"github.com/tigrisdata/todo/route"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := tigris.NewClient(ctx, &tigris.Config{Project: "todo"})

	if err != nil {
		panic(err)
	}
	defer client.Close()

	db, err := client.OpenDatabase(ctx,
		&model.Task{},
	)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	route.SetupTaskCRUD[model.Task](r, db, "tasks")

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
