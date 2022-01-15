package domain

import "context"

// CourseRepository defines the expected behaviour from a pkg storage.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
	FindAll(ctx context.Context) ([]Course, error)
	FindById(ctx context.Context, id string) (Course, error)
}

//go:generate mockery --case=snake --outpkg=mocks --output=storage/mocks --name=CourseRepository

// Course is the data structure that represents a pkg.
type Course struct {
	id          string
	name        string
	description string
}

// NewCourse creates a new pkg.
func NewCourse(id, name, description string) Course {
	return Course{
		id:          id,
		name:        name,
		description: description,
	}
}

// ID returns the pkg unique identifier.
func (c Course) ID() string {
	return c.id
}

// Name returns the pkg name.
func (c Course) Name() string {
	return c.name
}

// Description returns the pkg duration.
func (c Course) Description() string {
	return c.description
}
