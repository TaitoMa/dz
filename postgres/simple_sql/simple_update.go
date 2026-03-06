package simple_sql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = TRUE, completed_at = NOW()
	WHERE id = $1
`
	_, err := conn.Exec(ctx, sqlQuery, id)

	fmt.Println(id)

	return err
}

func UpdateTask(ctx context.Context, conn *pgx.Conn, task TaskModel) error {
	sqlQuery := `
	UPDATE tasks
	SET title=$1, description=$2, completed = $3, created_at=$4, completed_at = $5
	WHERE id = $6
`
	_, err := conn.Exec(ctx, sqlQuery, task.Title, task.Description, task.Completed, task.CreatedAt, task.CompletedAt, task.ID)

	return err
}
