package mysql

import (
	"database/sql"
	"strings"

	"faizisyellow.com/todolist/pkg/models"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(email, name, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (email, name, hashed_password, created_at)
	VALUES(?, ?, ?, NOW())`

	_, err = m.DB.Exec(stmt, email, name, string(hashedPassword))
	if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1062 && strings.Contains(mysqlError.Message, "users.email_UNIQUE") {
				return models.ErrDuplicateEmail
			}
		}
	}

	return err
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	row := m.DB.QueryRow("SELECT id, hashed_password FROM users WHERE email = ?", email)
	err := row.Scan(&id, &hashedPassword)
	// Check if the email exist.
	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	// Check if the password is match.
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	// And if, password is correct. Return the user ID
	return id, nil
}

func (m *UserModel) Get(id int) (*models.Users, error) {
	s := &models.Users{}

	stmt := `SELECT id, name, email, created_at FROM users WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	err := row.Scan(&s.ID, &s.Name, &s.Email, &s.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecords
	} else if err != nil {
		return nil, err
	}

	return s, nil
}
