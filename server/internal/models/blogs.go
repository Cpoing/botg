package models

import (
	"database/sql"
	"errors"
	"time"
)

type Blog struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type BlogModel struct {
	DB *sql.DB
}

func (m *BlogModel) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO blogs (title, content, created)
  VALUES(?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *BlogModel) Get(id int) (Blog, error) {
	stmt := `SELECT id, title, content, created FROM blogs
  WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	var b Blog

	err := row.Scan(&b.ID, &b.Title, &b.Content, &b.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Blog{}, ErrNoRecord
		} else {
			return Blog{}, err
		}
	}

	return b, nil
}

func (m *BlogModel) Latest() ([]Blog, error) {
	stmt := `SELECT id, title, content, created FROM blogs
  ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var blogs []Blog

	for rows.Next() {
		var b Blog

		err = rows.Scan(&b.ID, &b.Content, &b.Content, &b.Created)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, b)
	}

	return blogs, nil
}
