/*
Copyright belongs to all species of all the universes
*/
package cmd

import (
	"fmt"
	"github.com/mirantis/broker/cmd/container"
	"github.com/mirantis/broker/cmd/logs"
	"github.com/mirantis/broker/cmd/node"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "broker",
	Short: "Is a CLI tool to analyze/parse the dsinfo.json file and extract meaningful information",
	Run: func(cmd *cobra.Command, args []string) {
		help := `
A cli tool/parser/analyzer to analyze support bundle's dsinfo.json file and show meaningful output

Usage:  broker [OPTIONS] COMMAND

Available Commands:
  node	      Shows node specific information
  container   Shows container specific information 
  logs        Shows logs of a container or a specific object of a node
  info        Shows components related information of the nodes and container  
  stats	      Shows statistics/metrics/performance related information

Flags:
  -h, --help  Help for Broker
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
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(container.ContainerCmd)
	rootCmd.AddCommand(container.ContListCmd)
	rootCmd.AddCommand(container.ContInspectCmd)
	rootCmd.AddCommand(node.NodeCmd)
	rootCmd.AddCommand(logs.LogsCmd)
	rootCmd.AddCommand(container.ContListCmd)
	rootCmd.SetHelpCommand(helpCmd)

}
