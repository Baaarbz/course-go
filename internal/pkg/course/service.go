package course

import (
	domain "barbz.dev/course-go/internal"
	"context"
	"errors"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// Service is the interface that provides courses methods.
type Service interface {
	// CreateCourse save a new instance of course
	CreateCourse(ctx context.Context, id, name, duration string) error

	// FindAllCourses retrieve all registered courses
	FindAllCourses(ctx context.Context) ([]DTO, error)
}

type service struct {
	courseRepository domain.CourseRepository
}

// NewCourseService returns the default Service interface implementation.
func NewCourseService(courseRepository domain.CourseRepository) Service {
	return &service{
		courseRepository: courseRepository,
	}
}

// CreateCourse implements the course.Service interface.
func (s *service) CreateCourse(ctx context.Context, id, name, description string) error {
	if len(name) == 0 || len(description) == 0 {
		return ErrInvalidArgument
	}
	course := domain.NewCourse(id, name, description)
	return s.courseRepository.Save(ctx, course)
}

// FindAllCourses implements the course.Service interface.
func (s *service) FindAllCourses(ctx context.Context) ([]DTO, error) {
	courses, err := s.courseRepository.FindAll(ctx)
	return mapListCourseToDto(courses), err
}

type DTO struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func mapListCourseToDto(courses []domain.Course) []DTO {
	coursesDto := make([]DTO, len(courses))
	for index, element := range courses {
		// _ is the index where we are
		// element is the element from courses for where we are
		course := mapCourseToDto(element)
		coursesDto[index] = course
	}
	return coursesDto
}

func mapCourseToDto(course domain.Course) DTO {
	return DTO{
		ID:          course.ID(),
		Name:        course.Name(),
		Description: course.Description(),
	}
}
