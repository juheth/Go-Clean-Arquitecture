package users

import (
	"net/http"

	r "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	types "github.com/juheth/Go-Clean-Arquitecture/src/common/types"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/controllers"
	usecases "github.com/juheth/Go-Clean-Arquitecture/src/modules/usecases"
	"go.uber.org/fx"
)

func configureModuleRoutes(r *r.Result, h *types.HandlersStore, uc *controllers.UserController) {

	handlersModuleUsers := &types.SliceHandlers{
		Prefix: "users",
		Routes: []types.HandlerModule{
			{
				Route:   "/create",
				Method:  http.MethodPost,
				Handler: uc.CreateUser,
			},
			{
				Route:   "/all",
				Method:  http.MethodGet,
				Handler: uc.GetAllUsers,
			},
		},
	}
	h.Handlers = append(h.Handlers, *handlersModuleUsers)

}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(usecases.NewCreateUserUseCase),
		fx.Provide(controllers.NewUserController),

		fx.Invoke(configureModuleRoutes),
	}
}
