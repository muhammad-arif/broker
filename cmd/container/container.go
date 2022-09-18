/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/

package container

import (
	"fmt"
	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var ContainerCmd = &cobra.Command{
	Use:       "container",
	Short:     "Check containers information",
	ValidArgs: []string{"ls", "inspect"},
	Args:      cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("container called")
	},
	Aliases: []string{"c", "cont", "container"},
}

func init() {
	ContainerCmd.AddCommand(ContListCmd)
	ContainerCmd.AddCommand(contInspectCmd)
}
