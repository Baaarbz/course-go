package server

import (
	"barbz.dev/course-go/internal/pkg/course"
	"barbz.dev/course-go/internal/platform/server/handler/courses"
	"barbz.dev/course-go/internal/platform/server/handler/health"
	"barbz.dev/course-go/internal/platform/server/middleware/logging"
	"barbz.dev/course-go/internal/platform/server/middleware/recovery"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	engine          *gin.Engine
	httpAddr        string
	shutdownTimeout time.Duration
	// Dependencies
	courseService course.Service
}

func New(ctx context.Context, host string, port uint, courseService course.Service, shutdownTimeout time.Duration) (context.Context, Server) {
	engine := gin.New()
	// Register middlewares
	engine.Use(recovery.Middleware(), logging.Middleware())

	srv := Server{
		engine:          engine,
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		shutdownTimeout: shutdownTimeout,
		courseService:   courseService,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shutdown", err)
		}
	}()
	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutdown)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.courseService))
	s.engine.GET("/courses", courses.RetrieveAll(s.courseService))
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
