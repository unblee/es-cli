package delete

import (
	"context"

	delete "github.com/unblee/es-cli/cmd/delete/index"
	"github.com/unblee/es-cli/domain"
	"github.com/spf13/cobra"
)

func NewDeleteCommand(ctx context.Context, ind domain.Index) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete elasticsearch resources",
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(delete.NewIndexCmd(ctx, ind))
	return cmd
}
