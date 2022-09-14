/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all  EMEA/APAC/AMER TSE Colleagues
*/

package logs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/mirantis/broker/misc"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var containerlogCmd = &cobra.Command{
	Use:   "container [containerName] [nodeName]",
	Short: "Fetch the logs of the containers from specific node",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		getContainerLogs(args)
	},
}

func init() {
}

type nestedDsinfoType struct {
	ContainerInfo map[string]containerInfoT `json:"container_info"`
}
type containerInfoT struct {
	Logs json.RawMessage `json:"logs"`
}

func getContainerLogs(a []string) {
	cont := a[0]
	node := a[1]
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
		var NodeContents nestedDsinfoType
		err = sonic.Unmarshal(nodeDsinfoStruct.DsinfoContents, &NodeContents)
		if err != nil {
			fmt.Errorf("%v", err)
		}
		if d, isFound := NodeContents.ContainerInfo[cont]; isFound {
			//x, err := json.Marshal(d.Logs)
			if err != nil {
				fmt.Errorf("not unmarshalled properly, %v", err)
			}
			x := d.Logs
			//removing bracket
			x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
			x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9, 9}, []byte{})
			x = bytes.ReplaceAll(x, []byte{10, 34}, []byte{10})
			x = bytes.ReplaceAll(x, []byte{34, 44, 10}, []byte{10})
			x = bytes.ReplaceAll(x, []byte{92, 92, 92, 34}, []byte{32})
			x = bytes.ReplaceAll(x, []byte{92, 34}, []byte{34})
			x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 9, 9, 93}, []byte{})

			fmt.Println(string(x))
		} else {
			fmt.Errorf("container not found. Try with `broker ps -v`")
		}
	}
}
