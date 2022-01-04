package server

import (
	domain "barbz.dev/course-go/internal/platform"
	"barbz.dev/course-go/internal/platform/server/handler/courses"
	"barbz.dev/course-go/internal/platform/server/handler/health"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	engine           *gin.Engine
	httpAddr         string
	courseRepository domain.CourseRepository
}

func New(host string, port uint, courseRepository domain.CourseRepository) Server {
	srv := Server{
		engine:           gin.New(),
		httpAddr:         fmt.Sprintf("%s:%d", host, port),
		courseRepository: courseRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
}
