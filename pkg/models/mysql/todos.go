package mysql

import (
	"database/sql"
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
