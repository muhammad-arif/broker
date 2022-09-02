/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logs called")
	},
}
cd 
func init() {
	LogsCmd.AddCommand(nodelogCmd)
	LogsCmd.AddCommand(containerlogCmd)
}
