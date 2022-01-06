package courses

import (
	"barbz.dev/course-go/internal/pkg/course"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RetrieveAll(courseService course.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		courses, err := courseService.FindAllCourses(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		switch {
		case len(courses) == 0:
			ctx.Status(http.StatusNoContent)
			return
		default:
			ctx.JSON(http.StatusOK, courses)
			return
		}
	}
}
