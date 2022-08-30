/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package node

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NodeCmd represents the node command
var NodeCmd = &cobra.Command{
	Use:   "powerplug node [COMMAND]",
	Short: "Collect Swarm Node information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("node called")
	},
	Aliases: []string{"nodes", "n", "no"},
}

func init() {
	NodeCmd.AddCommand(nodeListCmd)
	NodeCmd.AddCommand(nodeInspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
