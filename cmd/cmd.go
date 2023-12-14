package cmd

import (
	"github.com/spf13/cobra"
)

// cmdCmd represents the run command
var cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "SSM commands",
	Long:  `SSM commands`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(cmdCmd)
}
