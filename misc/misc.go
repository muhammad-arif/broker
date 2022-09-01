package misc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fvbommel/sortorder"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"os"
	"sort"
)

//func ParseCoreDsinfo() []byte {
//	var file = "dsinfo.json"
//	dsinfoJson, err := os.ReadFile(file)
//	if err != nil {
//		fmt.Errorf("cannot Read File %s", err)
//	}
//	return dsinfoJson
//}

type ParseUcpNodesInspectT struct {
	UcpNodesTxt json.RawMessage `json:"ucp-nodes.txt"`
	//UcpNodesTxt json.RawMessage `json:"ucp-nodes.txt"`
}

// ParseUcpNodesInspect returns
// i) a slice of node names
// ii) node inspect structure
// iii) corefile (sharing os.ReadFile)
func ParseUcpNodesInspect() (map[string]string, *[]dLib.Node, *[]byte) {
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
	sort.Slice(nodeInspect, func(i, j int) bool {
		return sortorder.NaturalLess(string(nodeInspect[i].Spec.Role), string(nodeInspect[j].Spec.Role))
	})

	names := map[string]string{}
	for _, v := range nodeInspect {
		names[v.Description.Hostname] = v.ID
	}
	return names, &nodeInspect, &dsinfoJson
}

// PrettySTring will return pretty Json
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

/*
I need to test the perf of followng
1. read file then unmarshal whole struct vs unmarshal 1 struct
a=> One ReadFile = size of the file, unmarshaling is negligible
2. multiple small unmarshal and sharing readfile vs 1 big unmarshal and sharing objects
*/
