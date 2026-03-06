package simple_connection

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func CheckConnection() *pgx.Conn {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:1234@localhost:5432/postgres")
	defer conn.Close(ctx)
	if err != nil {
		panic(err)
	}

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Connection success!")
	return conn
}
