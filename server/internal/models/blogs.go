package models

import (
  "database/sql"
  "time"
)

type Blog struct {
  ID int
  Title string
  Content string
  Created time.Time
}

type BlogModel struct {
  DB *sql.DB
}

func (m *BlogModel) Insert(title string, content string) (int, error) {
  return 0, nil
}
func (m *BlogModel) Get(title string, content string) (Blog, error) {
  return Blog{}, nil
}

func (m *BlogModel) Latest() ([]Blog, error) {
  return nil, nil
}

