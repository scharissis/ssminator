package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the status of invoked commands",
	Long:  `Check the status of invoked commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdID, err := cmd.Flags().GetString("id")
		if err != nil || cmdID == "" {
			log.Fatalf("CheckCommandStatus - bad flag: %+v", err)
		}
		status, err := ssmi.CheckCommandStatus(cmdID)
		if err != nil {
			log.Fatalf("CheckCommandStatus failed: %+v", err)
		}
		jsonB, _ := json.MarshalIndent(status, "", " ")
		fmt.Println(string(jsonB))
	},
}

func init() {
	cmdCmd.AddCommand(checkCmd)

	checkCmd.Flags().String("id", "", "CommandID")
}
