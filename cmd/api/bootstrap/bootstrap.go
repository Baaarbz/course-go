package bootstrap

import (
	"barbz.dev/course-go/internal/pkg/course"
	"barbz.dev/course-go/internal/platform/server"
	"barbz.dev/course-go/internal/platform/storage/postgres"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "postgres"
	dbPass = "admin"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "course_go"
)

func Run() error {
	postgresURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", postgresURI)
	if err != nil {
		return err
	}

	courseRepository := postgres.NewCourseRepository(db)
	courseService := course.NewCourseService(courseRepository)

	srv := server.New(host, port, courseService)
	return srv.Run()
}
