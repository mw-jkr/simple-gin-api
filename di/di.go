//go:build wireinject

package di

import (
	"context"
	"github.com/google/wire"
)

func InitializeApplication(ctx context.Context) (*Application, error) {
	wire.Build(ApplicationSet)
	return &Application{}, nil
}
