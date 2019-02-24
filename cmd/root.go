package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var configFile string

func NewRootCmd() *cobra.Command {
	ctx := context.Background()
	cmd := &cobra.Command{
		Use:   "aegon",
		Short: "Aegon is an API Gateway",
		Long: `
		Lightweight API based on hellofresh/janus
		`,
	}

	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Config file (default is $PWD/aegon.toml)")
	// TODO: implement this
	cmd.AddCommand(NewCheckCmd(ctx))
	cmd.AddCommand(NewVersionCmd(ctx))
	// TODO: implement this
	cmd.AddCommand(NewServerStartCmd(ctx))

	return cmd
}
