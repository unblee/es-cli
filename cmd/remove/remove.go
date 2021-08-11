package remove

import (
	"context"

	"github.com/unblee/es-cli/cmd/remove/alias"
	"github.com/unblee/es-cli/domain"
	"github.com/spf13/cobra"
)

func NewRemoveCommand(ctx context.Context, alis domain.Alias) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove elasitcsearch resources",
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(alias.NewAliasCommand(ctx, alis))
	return cmd
}
