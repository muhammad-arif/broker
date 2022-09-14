/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/
package logs

import (
	"fmt"

	"github.com/spf13/cobra"
)

// LogsCmd represents the logs command
var LogsCmd = &cobra.Command{
	Use:   "logs [OPTIONS] [nodes| container] [nodeName|containerName nodeName]",
	Short: "Fetch the logs of the containers and nodes",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logs called")
	},
}

func init() {
	LogsCmd.AddCommand(nodelogCmd)
	LogsCmd.AddCommand(containerlogCmd)
}
