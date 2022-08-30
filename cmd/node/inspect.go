/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package node

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nodeInspectCmd represents the inspect command
var nodeInspectCmd = &cobra.Command{
	Use:   "powerplug node inspect [all|NODE-NAME]",
	Short: "Display detailed information on one or more nodes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inspect called")
	},
	Aliases: []string{"ins", "i", "inspect"},
}

func init() {
}
