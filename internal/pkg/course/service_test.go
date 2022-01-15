package course

import (
	"barbz.dev/course-go/internal/platform/storage/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_Service_CreateCourse_RepositoryError(t *testing.T) {
	course := DTO{
		ID:          "",
		Name:        "Test course",
		Description: "A long description of this test course.",
	}

	courseRepositoryMock := new(mocks.CourseRepository)
	courseRepositoryMock.
		On("Save", mock.Anything, mock.AnythingOfType("domain.Course")).
		Return(errors.New("test error"))

	courseService := NewCourseService(courseRepositoryMock)
	err := courseService.CreateCourse(context.Background(), course)

	courseRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_CreateCourse_InvalidArguments(t *testing.T) {
	course := DTO{
		ID:          "",
		Name:        "",
		Description: "A long description of this test course.",
	}

	courseRepositoryMock := new(mocks.CourseRepository)
	courseService := NewCourseService(courseRepositoryMock)

	err := courseService.CreateCourse(context.Background(), course)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidArgument)

	course.Name = "Test course"
	course.Description = ""
	err = courseService.CreateCourse(context.Background(), course)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidArgument)
}

func Test_Service_CreateCourse_Succeed(t *testing.T) {
	course := DTO{
		ID:          "",
		Name:        "Test course",
		Description: "A long description of this test course.",
	}

	courseRepositoryMock := new(mocks.CourseRepository)
	courseRepositoryMock.
		On("Save", mock.Anything, mock.AnythingOfType("domain.Course")).
		Return(nil)

	courseService := NewCourseService(courseRepositoryMock)
	err := courseService.CreateCourse(context.Background(), course)

	courseRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
