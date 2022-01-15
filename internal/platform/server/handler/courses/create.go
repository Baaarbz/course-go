package courses

import (
	"barbz.dev/course-go/internal/pkg/course"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(courseService course.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req course.DTO
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := courseService.CreateCourse(ctx, req); err != nil {
			switch {
			case errors.Is(err, course.ErrInvalidArgument):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
