/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/

package container

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/mirantis/broker/misc"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"github.com/spf13/cobra"
)

import (
	"encoding/json"
)

var pretty bool

// contInspectCmd represents the inspect command
var contInspectCmd = &cobra.Command{
	Use:   "inspect [CONTAINER NAME] [NODE NAME] ",
	Short: "Display detailed information on one or more containers",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		getContainerInspect(args)
	},
	Aliases: []string{"inspect", "ins", "describe"},
}

func init() {
	contInspectCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "pretty JSON output")
}

func getContainerInspect(a []string) {
	node := a[1]
	cont := a[0]
	// Invoking ParseUcpNodesInspect to collect the node names and core file
	nodeList, _, dsinfoJson := misc.ParseUcpNodesInspect()

	// checking if the mention node is present in the list
	if _, ok := nodeList[node]; ok {
		//If the node is valid
		// dumping core dsinfo structure to nestedDsinfo as a map so that the node can be parsed
		var nestedDsinfo = make(map[string]json.RawMessage)
		err := sonic.Unmarshal(*dsinfoJson, &nestedDsinfo)
		if err != nil {
			fmt.Errorf("cannot unmarshal %s", err)
		}
		var nodeDsinfoStruct dLib.DsinfoSlashDsinfoDotJson
		err = sonic.Unmarshal(nestedDsinfo[node], &nodeDsinfoStruct)
		if err != nil {
			fmt.Errorf("Cannot unmarshal nesteddsinfo")
		}
		var NodeContents NestedDsinfoT
		err = sonic.Unmarshal(nodeDsinfoStruct.DsinfoContents, &NodeContents)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		if d, isFound := NodeContents.ContainerInfo[cont]; isFound {
			x, err := json.Marshal(d.Inspect)
			if err != nil {
				fmt.Errorf("not unmarshalled properly, %v", err)
			}
			if pretty {
				z, _ := misc.PrettyString(string(x))
				fmt.Println(z)
			} else {
				fmt.Println(string(x))
			}
		} else {
			fmt.Errorf("container not found. Try with `broker ps -v`")
		}
	}
}
