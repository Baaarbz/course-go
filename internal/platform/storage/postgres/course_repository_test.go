package postgres

import (
	domain "barbz.dev/course-go/internal/platform"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CourseRepository_Save_Error(t *testing.T) {
	id, name, description := "357a6c62-94d5-4b22-9917-ca79d91867bc", "Test Course", "Long test course description."
	course := domain.NewCourse(id, name, description)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.
		ExpectExec("INSERT INTO course (name, description) VALUES ($1, $2)").
		WithArgs(name, description).
		WillReturnError(errors.New("something-failed"))

	repo := NewCourseRepository(db)
	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_CourseRepository_Save_Succeed(t *testing.T) {
	id, name, description := "357a6c62-94d5-4b22-9917-ca79d91867bc", "Test Course", "Long test course description."
	course := domain.NewCourse(id, name, description)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.
		ExpectExec("INSERT INTO course (name, description) VALUES ($1, $2)").
		WithArgs(name, description).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCourseRepository(db)
	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
