package usecase

import (
	"context"

	"github.com/israelalvesmelo/uow/internal/entity"
	"github.com/israelalvesmelo/uow/internal/repository"
	"github.com/israelalvesmelo/uow/pkg/uow"
)

type addCourseUseCaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUseCaseUow(uow uow.UowInterface) *addCourseUseCaseUow {
	return &addCourseUseCaseUow{
		Uow: uow,
	}
}

func (a *addCourseUseCaseUow) Execute(ctx context.Context, input InputUseCase) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		// tudo que estiver dentro do Do, vai ser executado dentro de uma transação
		category := entity.Category{
			Name: input.CategoryName,
		}
		repoCategory := a.getCategoryRepository(ctx)
		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:       input.CourseName,
			CategoryID: input.CourseCategoryID,
		}
		repoCourse := a.getCourseRepository(ctx)
		err = repoCourse.Insert(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *addCourseUseCaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CategoryRepository")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CategoryRepositoryInterface)
}

func (a *addCourseUseCaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "CourseRepository")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CourseRepositoryInterface)
}
