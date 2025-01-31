package di

import (
	"example/test-golang/config"
	"example/test-golang/controller"
	"example/test-golang/http"
	"example/test-golang/repository"
	"example/test-golang/service"

	"github.com/google/wire"
)

type Application struct {
	Srv *http.Server
}

var ApplicationSet = wire.NewSet(
	ConfigSet,
	ProviderSet,
	wire.Struct(new(Application), "*"),
)

var ConfigSet = wire.NewSet(
	config.ProvideConfig,
)

var ProviderSet = wire.NewSet(
	http.ProvideServer,
	controller.ProvideUserController,
	service.ProvideUserService,
	repository.ProvideUserRepository,
)
