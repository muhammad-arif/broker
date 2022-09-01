/*
Copyright belongs to all species of all the universes
*/
package node

import (
	"encoding/json"
	"fmt"
	"github.com/mirantis/powerplug/misc"
	"github.com/spf13/cobra"
)

var pretty bool

// nodeInspectCmd represents the inspect command
var nodeInspectCmd = &cobra.Command{
	Use:   "node inspect [NODE-NAME]",
	Short: "Display detailed information on one or more nodes",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(`Please provide a node name. To list the nodes you can use`, "\n\tpowerplug node ls")
			return
		} else if len(args) > 0 {
			GetInspect(args)
		}
	},
	Aliases: []string{"ins", "i", "inspect"},
}

func init() {
	nodeInspectCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "pretty JSON output")
}

func GetInspect(a []string) {
	nodeNames, nodeInspect, _ := misc.ParseUcpNodesInspect()
	for _, v := range a {
		if _, ok := nodeNames[v]; ok {
			for _, inspect := range *nodeInspect {
				if inspect.Description.Hostname != v {
					continue
				} else {
					inspectInJson, _ := json.Marshal(inspect)
					if pretty {
						z, _ := misc.PrettyString(string(inspectInJson))
						fmt.Println(z)
					} else {
						fmt.Println(string(inspectInJson))
					}
				}
			}
		} else {
			fmt.Println("Invalid  Node")
			return
		}
	}
}
