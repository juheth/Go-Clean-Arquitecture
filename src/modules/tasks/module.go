package tasks

import (
	"net/http"

	r "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	types "github.com/juheth/Go-Clean-Arquitecture/src/common/types"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/tasks/controllers"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/tasks/domain/repository"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/tasks/usecases"
	"go.uber.org/fx"
)

func configureModuleRoutes(r *r.Result, h *types.HandlersStore, uc *controllers.TaskController) {

	handlersModuleTasks := &types.SliceHandlers{

		Prefix: "tasks",
		Routes: []types.HandlerModule{
			{
				Route:   "/create",
				Method:  http.MethodPost,
				Handler: uc.CreateTask,
			},
			{
				Route:   "/all",
				Method:  http.MethodGet,
				Handler: uc.GetAllTasks,
			},
			{
				Route:   "/:id",
				Method:  http.MethodGet,
				Handler: uc.GetTaskByID,
			},
			{
				Route:   "/update/:id",
				Method:  http.MethodPut,
				Handler: uc.UpdateTask,
			},
			{
				Route:   "/delete/:id",
				Method:  http.MethodDelete,
				Handler: uc.DeleteTask,
			},
			{
				Route:   "/update-status/:id",
				Method:  http.MethodPut,
				Handler: uc.UpdateTaskStatus,
			},
			{
				Route:   "/project/:project_id/tasks",
				Method:  http.MethodGet,
				Handler: uc.GetTasksByProjectID,
			},
		},
	}

	h.Handlers = append(h.Handlers, *handlersModuleTasks)
}

func ModuleProviders() []fx.Option {

	return []fx.Option{
		fx.Provide(repository.NewTaskRepository),
		fx.Provide(usecases.NewTaskUseCase),

		fx.Provide(controllers.NewTaskController),

		fx.Invoke(configureModuleRoutes),
	}
}
