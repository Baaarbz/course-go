package bootstrap

import (
	"barbz.dev/course-go/internal/pkg/course"
	"barbz.dev/course-go/internal/platform/server"
	"barbz.dev/course-go/internal/platform/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser = "postgres"
	dbPass = "admin"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "course_go"
)

func Run() error {
	db, err := initDatabase()
	if err != nil {
		return err
	}

	courseRepository := postgres.NewCourseRepository(db)
	courseService := course.NewCourseService(courseRepository)

	ctx, srv := server.New(context.Background(), host, port, courseService, shutdownTimeout)
	return srv.Run(ctx)
}

func initDatabase() (*sql.DB, error) {
	postgresURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return nil, err
	}
	// Ping DB to check if the connection was established successfully
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
