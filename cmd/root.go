package cmd

import (
	//"context"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	// ctx := context.Background()
	cmd := &cobra.Command{
		Use:   `aegon`,
		Short: `aegon is an API Gateway`,
		Long: `
Use Aegon for your API Gateway. Documentation @github.com/alamin-mahamud/aegon
		`,
	}
	return cmd
}
