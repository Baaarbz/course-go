// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "barbz.dev/course-go/internal/platform"
	mock "github.com/stretchr/testify/mock"
)

// CourseRepository is an autogenerated mock type for the CourseRepository type
type CourseRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, course
func (_m *CourseRepository) Save(ctx context.Context, course domain.Course) error {
	ret := _m.Called(ctx, course)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Course) error); ok {
		r0 = rf(ctx, course)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
