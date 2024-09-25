package repository

import (
	"context"
	"database/sql"

	"github.com/israelalvesmelo/uow/internal/db"
	"github.com/israelalvesmelo/uow/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type courseRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(dbt *sql.DB) *courseRepository {
	return &courseRepository{
		DB:      dbt,
		Queries: db.New(dbt),
	}
}

func (r *courseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCourse(ctx, db.CreateCourseParams{
		Name:       course.Name,
		CategoryID: int32(course.CategoryID),
	})
}
