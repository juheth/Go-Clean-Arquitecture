package usecases

import (
	"net/http"

	r "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	types "github.com/juheth/Go-Clean-Arquitecture/src/common/types"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/projects/controllers"

	"github.com/juheth/Go-Clean-Arquitecture/src/modules/projects/domain/repository"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/projects/usecases"
	"go.uber.org/fx"
)

func configureModuleRoutes(r *r.Result, h *types.HandlersStore, uc *controllers.ProjectController) {

	handlersModuleProject := &types.SliceHandlers{

		Prefix: "projects",
		Routes: []types.HandlerModule{
			{
				Route:   "/create",
				Method:  http.MethodPost,
				Handler: uc.CreateProject,
			},
			{
				Route:   "/all",
				Method:  http.MethodGet,
				Handler: uc.GetAllProjects,
			},
			{
				Route:   "/:id",
				Method:  http.MethodGet,
				Handler: uc.GetProjectByID,
			},
			{
				Route:   "/update/:id",
				Method:  http.MethodPut,
				Handler: uc.UpdateProject,
			},
			{
				Route:   "/delete/:id",
				Method:  http.MethodDelete,
				Handler: uc.DeleteProject,
			},
			// {
			// Route:   "/update-status/:id",
			// Method:  http.MethodPut,
			// 	Handler: uc.UpdateProjectStatus,
			// },
		},
	}

	h.Handlers = append(h.Handlers, *handlersModuleProject)
}

func ModuleProviders() []fx.Option {

	return []fx.Option{
		fx.Provide(repository.NewProjectRepository),
		fx.Provide(usecases.NewProjectUseCase),
		fx.Provide(controllers.NewProjectController),

		fx.Invoke(configureModuleRoutes),
	}
}
