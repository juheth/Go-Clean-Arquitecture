package users

import (
	"net/http"

	r "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	types "github.com/juheth/Go-Clean-Arquitecture/src/common/types"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/controllers"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/repository"
	usecases "github.com/juheth/Go-Clean-Arquitecture/src/modules/users/usecases"
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
			{
				Route:   "/:id",
				Method:  http.MethodGet,
				Handler: uc.GetUserByID,
			},
			{
				Route:   "/authenticated",
				Method:  http.MethodPost,
				Handler: uc.JwtUser,
			},
			{
				Route:   "/check-auth",
				Method:  http.MethodPost,
				Handler: uc.IsAuthenticated,
			},
			{
				Route:   "/update/:id",
				Method:  http.MethodPut,
				Handler: uc.UpdateUser,
			},
			{
				Route:   "/delete/:id",
				Method:  http.MethodDelete,
				Handler: uc.DeleteUser,
			},
		},
	}

	h.Handlers = append(h.Handlers, *handlersModuleUsers)
}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(repository.NewUserRepository),
		fx.Provide(usecases.NewUserUseCase),
		fx.Provide(controllers.NewUserController),

		fx.Invoke(configureModuleRoutes),
	}
}
