/*
Copyright belongs to all species of all the universes
*/
package container

import (
	"fmt"
	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var ContainerCmd = &cobra.Command{
	Use:   "container COMMAND",
	Short: "Check containers information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("container called")
	},
	Aliases: []string{"c", "cont", "container"},
}

func init() {
	ContainerCmd.AddCommand(ContListCmd)
	ContainerCmd.AddCommand(ContInspectCmd)
}
