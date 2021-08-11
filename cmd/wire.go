//+build wireinject

package cmd

import (
	"context"

	"github.com/google/wire"
	"github.com/unblee/es-cli/config"
	"github.com/unblee/es-cli/domain"
	"github.com/unblee/es-cli/infra/es"
	"github.com/unblee/es-cli/infra/http"
	"github.com/unblee/es-cli/infra/logger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func InitializeCmd(ctx context.Context, cfg config.Config) (*cobra.Command, error) {
	wire.Build(NewCmdRoot, es.NewBaseClient, http.NewClient, domain.NewIndex, domain.NewDetail, domain.NewAlias)
	return &cobra.Command{}, nil
}

func InitializeLogger(cfg config.Config) (*zap.Logger, error) {
	wire.Build(logger.NewLogger)
	return &zap.Logger{}, nil
}
