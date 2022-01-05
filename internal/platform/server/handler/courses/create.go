package courses

import (
	"barbz.dev/course-go/internal/pkg/course"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(courseService course.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := courseService.CreateCourse(ctx, req.ID, req.Name, req.Description); err != nil {
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
