package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name string, description string, categoryId string) (Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec(
		"INSERT INTO courses(id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryId,
	)
	if err != nil {
		return Course{}, err
	}

	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryId,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return buildCourses(rows)
}

func (c *Course) FindaByCategoryId(categoryId string) ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses WHERE category_id = $1", categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return buildCourses(rows)
}

func buildCourses(rows *sql.Rows) ([]Course, error) {
	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}

		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryId})
	}

	return courses, nil
}
