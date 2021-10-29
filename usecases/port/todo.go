package port

import (
	"context"
	"go-cla-mysql/entities/model"
	"go-cla-mysql/usecases/dto"
)

type TodoInputPort interface {
	Create(todo *model.Todo) (bool, error)
	FindAll(ctx context.Context, max int) (*dto.TodoOutPutUseCaseDto, error)
	FindByID(ctx context.Context, id int) (*dto.TodoOutPutUseCaseDto, error)
	Update(ctx context.Context, todo model.Todo) (model.Todos, error)
}

type TodoOutputPort interface {
	Render(todos *model.Todos) dto.TodoOutPutUseCaseDto
	RenderError(error)
}
