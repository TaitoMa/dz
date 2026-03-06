package main

import (
	"context"
	"demo/app-1/postgres/simple_connection"
	"demo/app-1/postgres/simple_sql"
	"fmt"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	err = simple_sql.CreateTable(ctx, conn)
	if err != nil {
		panic(err)
	}

	//if err = simple_sql.InsertRow(
	//	ctx,
	//	conn,
	//	"Проверить работу $",
	//	"Передать аргументы через $",
	//	false,
	//	time.Now(),
	//); err != nil {
	//	panic(err)
	//}

	//if err = simple_sql.UpdateRow(
	//	ctx,
	//	conn,
	//	3,
	//); err != nil {
	//	panic(err)
	//}

	tasks, err := simple_sql.SelectRow(
		ctx,
		conn,
	)
	if err != nil {
		panic(err)
	}

	for _, task := range tasks {
		fmt.Println("....................................................")
		fmt.Println(task)
		fmt.Println("....................................................")
	}

	//now := time.Now()
	//
	//newTask := simple_sql.TaskModel{
	//	ID:          8,
	//	Title:       "Новая изменённая таска",
	//	Description: "Изменение",
	//	Completed:   true,
	//	CreatedAt:   time.Now(),
	//	CompletedAt: &now,
	//}
	//
	//err = simple_sql.UpdateTask(
	//	ctx,
	//	conn,
	//	newTask,
	//)
	//if err != nil {
	//	panic(err)
	//}

	//if err = simple_sql.DeleteRow(ctx, conn); err != nil {
	//	panic(err)
	//}

	fmt.Println("Success")
}

// миграция
// создание up и down файлов migrate create -ext sql -dir migrations -seq init
// migrate -path migrations -database "postgres://postgres:1234@localhost:5432/postgres?sslmode=disable" up
