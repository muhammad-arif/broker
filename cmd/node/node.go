/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/

package node

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NodeCmd represents the node command
var NodeCmd = &cobra.Command{
	Use:   "node [COMMAND]",
	Short: "Collect Swarm Node information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node called")
	},
	Aliases: []string{"nodes", "n", "no", "node"},
}

func init() {
	NodeCmd.AddCommand(nodeListCmd)
	NodeCmd.AddCommand(nodeInspectCmd)
}
