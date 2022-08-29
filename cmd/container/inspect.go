/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var contInspectCmd = &cobra.Command{
	Use:   "powerplug container inspect [NODE NAME] [CONTAINER NAME]",
	Short: "Display detailed information on one or more containers",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inspect called")
	},
}

func init() {
	//ContainerCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
