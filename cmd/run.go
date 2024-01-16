package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run commands",
	Long:  `Run commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdInput, err := ssmi.DefaultCmdInput()
		if err != nil {
			log.Fatal("failed", err)
		}
		out, err := ssmi.RunOnAllInstances(cmdInput)
		if err != nil {
			log.Fatal("failed", err)
		}
		log.Printf("out:\n%+v\n", out)
	},
}

func init() {
	cmdCmd.AddCommand(runCmd)
}
