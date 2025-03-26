package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	"faizisyellow.com/todolist/pkg/models"
	"github.com/go-sql-driver/mysql"
)

type TodosModel struct {
	DB *sql.DB
}

func (t *TodosModel) Insert(task string, userID int) error {
	stmt := `INSERT INTO todos (task, user_id, created_at) VALUES(?, ?, NOW())`

	_, err := t.DB.Exec(stmt, task, userID)
	if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1452 && strings.Contains(mysqlError.Message, "a foreign key constraint fails") {
				return models.ErrRequireUser
			}
		}
	}

	return err
}

func (t *TodosModel) Latest(userID int) ([]*models.Todos, error) {
	stmt := `SELECT id, task, status, created_at FROM todos WHERE user_id = ? AND DATE(created_at) = CURDATE()`

	row, err := t.DB.Query(stmt, userID)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var todos []*models.Todos

	for row.Next() {
		t := &models.Todos{}

		err := row.Scan(&t.ID, &t.Task, &t.Status, &t.CreatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, t)
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *TodosModel) Get(id int) (*models.Todos, error) {

	todo := &models.Todos{}

	stmt := `SELECT id, task, status, created_at, user_id FROM todos WHERE id = ?`

	row := t.DB.QueryRow(stmt, id)
	err := row.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.CreatedAt, &todo.UserID)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecords
	} else if err != nil {
		return nil, err
	}

	return todo, nil

}

func (t *TodosModel) Update(col string, val string, id int) error {
	stmt := fmt.Sprintf("UPDATE todos SET %s = ? WHERE id = ?", col)

	_, err := t.DB.Exec(stmt, val, id)
	if err != nil {
		return err
	}

	return nil
}
