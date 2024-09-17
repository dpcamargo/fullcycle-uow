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

func (c *Course) Create(name, description, categoryID string) (*Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?)",
		id, name, description, categoryID)
	if err != nil {
		return nil, err
	}
	return &Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	courses, err := c.db.Query("SELECT id, name, description, category_id FROM courses")
	if err != nil {
		return nil, err
	}
	defer courses.Close()

	var coursesReturn []Course
	for courses.Next() {
		var id, name, description, categoryID string
		if err = courses.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		coursesReturn = append(coursesReturn, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return coursesReturn, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	course, err := c.db.Query("SELECT id, name, description, category_id FROM courses WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer course.Close()

	var courses []Course
	for course.Next() {
		var id, name, description, categoryID string
		if err = course.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}
		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryID: categoryID})
	}
	return courses, nil
}
