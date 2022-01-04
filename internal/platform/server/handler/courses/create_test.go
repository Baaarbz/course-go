package courses

import (
	"barbz.dev/course-go/internal/platform/storage/mocks"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler_Create(t *testing.T) {
	courseRepository := new(mocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("domain.Course")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", CreateHandler(courseRepository))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:   "357a6c62-94d5-4b22-9917-ca79d91867bc",
			Name: "Test Course",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, res.StatusCode, http.StatusBadRequest)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:          "357a6c62-94d5-4b22-9917-ca79d91867bc",
			Name:        "Test Course",
			Description: "A long description of this test course.",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, res.StatusCode, http.StatusCreated)
	})
}
