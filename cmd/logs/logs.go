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
	Use:       "logs [node|container] [nodeName|containerName]",
	Short:     "Fetch the logs of the containers and nodes",
	ValidArgs: []string{"node", "container"},
	Args:      cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logs called")
	},
}

func init() {
	LogsCmd.AddCommand(nodelogCmd)
	LogsCmd.AddCommand(containerlogCmd)
}
