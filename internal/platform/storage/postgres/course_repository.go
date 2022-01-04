package postgres

import (
	domain "barbz.dev/course-go/internal/platform"
	"context"
	"database/sql"
	"fmt"
)

const (
	insert      = "insert into $1 (name, duration) values ($2, $3)"
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
	_, err := r.db.ExecContext(ctx, insert, courseTable, course.Name(), course.Duration())
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}
	return nil
}
