/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Command to show the use of broker commands",
	Run: func(cmd *cobra.Command, args []string) {
		help := `
A cli tool parser/analyzer to analyze support bundle's dsinfo.json file and show meaningful output

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

func init() {
}
