package get

import (
	"context"

	get "github.com/unblee/es-cli/cmd/get/detai"
	"github.com/unblee/es-cli/domain"
	"github.com/spf13/cobra"
)

func NewGetCommand(ctx context.Context, dtl domain.Detail) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get up elasitcsearch resources",
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(get.NewDetailCmd(ctx, dtl))
	return cmd
}
