package models

import (
	"database/sql"
	"errors"
)

type ProgressModelInterface interface {
	Insert(user_id int, last_lesson, last_position string) (int, error)
	Get(id int) (*Progress, error)
	Update(last_lesson, last_position string, id int) (*Progress, error)
}

type Progress struct {
	UserID       int
	LastLesson   string
	LastPosition string
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
type ProgressModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *ProgressModel) Insert(user_id int, last_lesson string, last_position string) (int, error) {

	stmt := `INSERT INTO progress (user_id, last_lesson, last_position) VALUES($1, $2, $3)`

	result, err := m.DB.Exec(stmt, user_id, last_lesson, last_position)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

// This will return a specific snippet based on its id.
func (m *ProgressModel) Get(id int) (*Progress, error) {

	stmt := `SELECT user_id, last_lesson, last_position FROM progress
	WHERE user_id = $1`

	p := &Progress{}

	err := m.DB.QueryRow(stmt, id).Scan(&p.UserID, &p.LastLesson, &p.LastPosition)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return p, nil
}

func (m *ProgressModel) Update(last_lesson string, last_position string, id int) (*Progress, error) {

	stmt := `UPDATE progress SET last_lesson = $1, last_position = $2  WHERE user_id = $3`

	p := &Progress{}

	err := m.DB.QueryRow(stmt, last_lesson, last_position, id).Scan(&p.UserID, &p.LastLesson, &p.LastPosition)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return p, nil
}
