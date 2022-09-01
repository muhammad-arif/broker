/*
Copyright belongs to all species of all the universes
*/
package cmd

import (
	"fmt"
	"github.com/mirantis/powerplug/cmd/container"
	"github.com/mirantis/powerplug/cmd/logs"
	"github.com/mirantis/powerplug/cmd/node"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "broker",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		help := `
Usage:  broker [OPTIONS] COMMAND

A parser/analyzer to analyze support bundle's dsinfo.json file and show meaningful output
  
  node	      Shows node specific information
  contianer   Shows container specific information 
  logs        Shows logs of a container or a specific object of a node
  info        Shows components related information of the nodes and container  
  stats	      Shows statistics/metrics/performance related information

`
		fmt.Println(help)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(container.ContainerCmd)
	rootCmd.AddCommand(container.ContListCmd)
	rootCmd.AddCommand(container.ContInspectCmd)
	rootCmd.AddCommand(node.NodeCmd)
	rootCmd.AddCommand(logs.LogsCmd)
	rootCmd.AddCommand(container.ContListCmd)
}
