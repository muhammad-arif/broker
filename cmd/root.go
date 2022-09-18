/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/
package cmd

import (
	"github.com/mirantis/broker/cmd/container"
	"github.com/mirantis/broker/cmd/logs"
	"github.com/mirantis/broker/cmd/node"
	"github.com/spf13/cobra"
	"os"
)

var help bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:       "broker",
	Short:     "A cli tool/parser/analyzer to analyze support bundle's dsinfo.json file and show meaningful output",
	Args:      cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	ValidArgs: []string{"container", "node", "log", "info", "stats", "network"},
	Run: func(cmd *cobra.Command, args []string) {
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
	rootCmd.AddCommand(node.NodeCmd)
	rootCmd.AddCommand(logs.LogsCmd)
	rootCmd.SetHelpCommand(helpCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(networkCmd)

	// Flags for top level commands (root, info, network ...)
	rootCmd.Flags().BoolVarP(&help, "help", "h", false, "help command")
}
