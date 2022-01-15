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
	CreateCourse(ctx context.Context, course DTO) error

	// FindAllCourses retrieve all registered courses
	FindAllCourses(ctx context.Context) ([]DTO, error)

	// FindCourse retrieve course searching by id
	FindCourse(ctx context.Context, id string) (DTO, error)
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
func (s *service) CreateCourse(ctx context.Context, courseDto DTO) error {
	if len(courseDto.Name) == 0 || len(courseDto.Description) == 0 {
		return ErrInvalidArgument
	}
	course := mapDtoToCourse(courseDto)
	return s.courseRepository.Save(ctx, course)
}

// FindAllCourses implements the course.Service interface.
func (s *service) FindAllCourses(ctx context.Context) ([]DTO, error) {
	courses, err := s.courseRepository.FindAll(ctx)
	return mapListCourseToDto(courses), err
}

// FindCourse implements the course.Service interface.
func (s *service) FindCourse(ctx context.Context, id string) (DTO, error) {
	course, err := s.courseRepository.FindById(ctx, id)
	return mapCourseToDto(course), err
}
