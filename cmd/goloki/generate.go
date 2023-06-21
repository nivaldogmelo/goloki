package goloki

import (
	"github.com/nivaldogmelo/goloki/pkg/goloki"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"gen"},
	Short:   "Generate logs",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		goloki.Generate(args[0])
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
