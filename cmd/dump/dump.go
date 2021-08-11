package dump

import (
	"context"

	dump "github.com/unblee/es-cli/cmd/dump/index"
	"github.com/unblee/es-cli/domain"
	"github.com/spf13/cobra"
)

func NewDumpCommand(ctx context.Context, ind domain.Index) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dump",
		Short: "Delete elasticsearch resources",
		Args:  cobra.ExactArgs(2),
	}

	cmd.AddCommand(dump.NewIndexCmd(ctx, ind))
	return cmd
}
