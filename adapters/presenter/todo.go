package presenter

import (
	"go-cla-mysql/entities/model"
	"go-cla-mysql/usecases/dto"
	"go-cla-mysql/usecases/port"
	"net/http"
)

type Todo struct {
	w http.ResponseWriter
}

func NewTodoOputPutPort(w http.ResponseWriter) port.TodoOutputPort {
	return &Todo{
		w: w,
	}
}

// usecasesのTodoOutputPortを実装
func (t *Todo) Render(todos *model.Todos) dto.TodoOutPutUseCaseDto {
	t.w.WriteHeader(http.StatusOK)
	var output dto.TodoOutPutUseCaseDto
	output.Hits = len(*todos)
	output.Todos = *todos
	return output
}

func (t *Todo) RenderError(err error) {
	panic("implement me")
}