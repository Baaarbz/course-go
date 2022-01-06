package server

import (
	"barbz.dev/course-go/internal/pkg/course"
	"barbz.dev/course-go/internal/platform/server/handler/courses"
	"barbz.dev/course-go/internal/platform/server/handler/health"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	engine        *gin.Engine
	httpAddr      string
	courseService course.Service
}

func New(host string, port uint, courseService course.Service) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		// Dependencies
		courseService: courseService,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.courseService))
	s.engine.GET("/courses", courses.RetrieveAll(s.courseService))
}
