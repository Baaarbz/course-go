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
	courseId := ""
	courseName := "Test course"
	courseDescription := "A long description of this test course."

	courseRepositoryMock := new(mocks.CourseRepository)
	courseRepositoryMock.
		On("Save", mock.Anything, mock.AnythingOfType("domain.Course")).
		Return(errors.New("test error"))

	courseService := NewCourseService(courseRepositoryMock)
	err := courseService.CreateCourse(context.Background(), courseId, courseName, courseDescription)

	courseRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_Service_CreateCourse_InvalidArguments(t *testing.T) {
	courseId := ""
	courseName := ""
	courseDescription := "A long description of this test course."

	courseRepositoryMock := new(mocks.CourseRepository)
	courseService := NewCourseService(courseRepositoryMock)

	err := courseService.CreateCourse(context.Background(), courseId, courseName, courseDescription)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidArgument)

	courseName = "Test course"
	courseDescription = ""
	err = courseService.CreateCourse(context.Background(), courseId, courseName, courseDescription)

	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrInvalidArgument)
}

func Test_Service_CreateCourse_Succeed(t *testing.T) {
	courseId := ""
	courseName := "Test course"
	courseDescription := "A long description of this test course."

	courseRepositoryMock := new(mocks.CourseRepository)
	courseRepositoryMock.
		On("Save", mock.Anything, mock.AnythingOfType("domain.Course")).
		Return(nil)

	courseService := NewCourseService(courseRepositoryMock)
	err := courseService.CreateCourse(context.Background(), courseId, courseName, courseDescription)

	courseRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
