package postgres

import (
	domain "barbz.dev/course-go/internal"
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

const (
	courseTable = "course"
)

// CourseRepository is a Postgres domain.CourseRepository implementation.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a Postgres-based implementation of domain.CourseRepository
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the domain.CourseRepository
func (r *CourseRepository) Save(ctx context.Context, course domain.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlBuilderCreateCourse))
	query, args := courseSQLStruct.InsertInto(courseTable, sqlBuilderCreateCourse{
		Name:        course.Name(),
		Description: course.Description(),
	}).BuildWithFlavor(sqlbuilder.PostgreSQL)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist pkg on database: %v", err)
	}
	return nil
}

// FindAll implements the domain.CourseRepository
func (r *CourseRepository) FindAll(ctx context.Context) ([]domain.Course, error) {
	var courses []domain.Course

	courseSQLStruct := sqlbuilder.NewStruct(new(sqlBuilderCourse))
	query, _ := courseSQLStruct.SelectFrom(courseTable).BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return courses, fmt.Errorf("error trying to select all courses from database: %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var course domain.Course
		if err := rows.Scan(&course); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			return courses, fmt.Errorf("error trying to scan row results of courses: %v", err)
		}
		courses = append(courses, course)
	}
	return courses, nil
}
