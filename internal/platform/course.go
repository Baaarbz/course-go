package domain

import "context"

// CourseRepository defines the expected behaviour from a course storage.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=mocks --output=storage/mocks --name=CourseRepository

// Course is the data structure that represents a course.
type Course struct {
	id          string
	name        string
	description string
}

// NewCourse creates a new course.
func NewCourse(id, name, description string) Course {
	return Course{
		id:          id,
		name:        name,
		description: description,
	}
}

// ID returns the course unique identifier.
func (c Course) ID() string {
	return c.id
}

// Name returns the course name.
func (c Course) Name() string {
	return c.name
}

// Description returns the course duration.
func (c Course) Description() string {
	return c.description
}
