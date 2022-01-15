package bootstrap

import (
	"barbz.dev/course-go/internal/pkg/course"
	"barbz.dev/course-go/internal/platform/server"
	"barbz.dev/course-go/internal/platform/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"time"
)

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s" split_words:"true"`
	// Database configuration
	DbUser    string        `required:"true" split_words:"true"`
	DbPass    string        `required:"true" split_words:"true"`
	DbHost    string        `required:"true" split_words:"true"`
	DbPort    string        `required:"true" split_words:"true"`
	DbName    string        `required:"true" split_words:"true"`
	DbTimeout time.Duration `default:"5s" split_words:"true"`
}

var cfg config

func Run() error {
	// Load config
	if err := envconfig.Process("course_go", &cfg); err != nil {
		return err
	}

	db, err := initDatabase()
	if err != nil {
		return err
	}

	courseRepository := postgres.NewCourseRepository(db, cfg.DbTimeout)
	courseService := course.NewCourseService(courseRepository)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, courseService, cfg.ShutdownTimeout)
	return srv.Run(ctx)
}

func initDatabase() (*sql.DB, error) {
	postgresURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DbUser,
		cfg.DbPass,
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbName)

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
