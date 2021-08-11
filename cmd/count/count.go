package count

import (
	"context"

	count "github.com/unblee/es-cli/cmd/count/index"
	"github.com/unblee/es-cli/domain"
	"github.com/spf13/cobra"
)

func NewCountCommand(ctx context.Context, ind domain.Index) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "count",
		Short: "Count elasticsearch resources",
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(count.NewIndexCmd(ctx, ind))
	return cmd
}
