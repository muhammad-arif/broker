/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package logs

import (
	"fmt"

	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var containerlogCmd = &cobra.Command{
	Use:   "container [containerName] [nodeName]",
	Short: "Fetch the logs of the containers from specific node",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("container called")
	},
}

func init() {
}
