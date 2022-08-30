package misc

import (
	"encoding/json"
	"fmt"
	"github.com/fvbommel/sortorder"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"os"
	"sort"
)

type ParseUcpNodesInspectT struct {
	UcpNodesTxt json.RawMessage `json:"ucp-nodes.txt"`
	//UcpNodesTxt json.RawMessage `json:"ucp-nodes.txt"`
}

// ParseUcpNodesInspect returns
// i) a slice of node names
// ii) node inspect structure
// iii) corefile (sharing os.ReadFile)
func ParseUcpNodesInspect() ([]string, *[]dLib.Node, *[]byte) {
	// reading file
	var file = "dsinfo.json"
	dsinfoJson, err := os.ReadFile(file)
	if err != nil {
		fmt.Errorf("cannot Read File %s", err)
	}
	// Unmarshalling to mininmal type ParseUcpNodesInspectType
	var ucpNodesInspect ParseUcpNodesInspectT
	err = json.Unmarshal(dsinfoJson, &ucpNodesInspect)
	if err != nil {
		fmt.Errorf("Cannot unmarshal %s", err)
	}
	// Creating a slice of node's hostname
	nodeInspect := dLib.NewNodeParser().GetAll(&ucpNodesInspect.UcpNodesTxt)
	names := []string{}
	for _, v := range nodeInspect {
		names = append(names, v.Description.Hostname)
	}
	sort.Slice(nodeInspect, func(i, j int) bool {
		return sortorder.NaturalLess(string(nodeInspect[i].Spec.Role), string(nodeInspect[j].Spec.Role))
	})
	return names, &nodeInspect, &dsinfoJson
}

/*
I need to test the perf of followng
1. read file then unmarshal whole struct vs unmarshal 1 struct
a=> One ReadFile = size of the file, unmarshaling is negligible
2. multiple small unmarshal and sharing readfile vs 1 big unmarshal and sharing objects
*/
