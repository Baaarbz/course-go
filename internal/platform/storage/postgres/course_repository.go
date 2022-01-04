package postgres

import (
	domain "barbz.dev/course-go/internal/platform"
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
)

const (
	courseTable = "courses"
)

// CourseRepository is a Postgres platform.CourseRepository implementation.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a Postgres-based implementation of platform.CourseRepository
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the platform.CourseRepository
func (r *CourseRepository) Save(ctx context.Context, course domain.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlBuilderCourse))
	query, args := courseSQLStruct.InsertInto(courseTable, sqlBuilderCourse{
		ID:          course.ID(),
		Name:        course.Name(),
		Description: course.Description(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}
	return nil
}
