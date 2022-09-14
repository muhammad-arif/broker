/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/

package misc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/fvbommel/sortorder"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"os"
	"sort"
	"strings"
)

const shortLen = 12

type ImageMetadata struct {
	Containers   string `json:"Containers"`
	CreatedAt    string `json:"CreatedAt"`
	CreatedSince string `json:"CreatedSince"`
	Digest       string `json:"Digest"`
	ID           string `json:"ID"`
	Repository   string `json:"Repository"`
	SharedSize   string `json:"SharedSize"`
	Size         string `json:"Size"`
	Tag          string `json:"Tag"`
	UniqueSize   string `json:"UniqueSize"`
	VirtualSize  string `json:"VirtualSize"`
}

func TruncateID(id string) string {
	if i := strings.IndexRune(id, ':'); i >= 0 {
		id = id[i+1:]
	}
	if len(id) > shortLen {
		id = id[:shortLen]
	}
	return id
}

type ParseUcpNodesInspectT struct {
	UcpNodesTxt json.RawMessage `json:"ucp-nodes.txt"`
}

// ParseUcpNodesInspect returns
// i) a slice of node names
// ii) node inspect structure
// iii) corefile (sharing os.ReadFile)
func ParseUcpNodesInspect() (map[string]string, *[]dLib.Node, *[]byte) {
	//defer profile.Start(profile.MemProfileHeap).Stop()
	//defer profile.Start(profile.MemProfileAllocs).Stop()
	// reading file
	var file = "dsinfo.json"
	dsinfoJson, err := os.ReadFile(file)
	if err != nil {
		fmt.Errorf("cannot Read File %s", err)
	}
	// Unmarshalling to mininmal type ParseUcpNodesInspectType
	var ucpNodesInspect ParseUcpNodesInspectT
	err = sonic.Unmarshal(dsinfoJson, &ucpNodesInspect)
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

func PrettyPrintFromByteSlice(d *json.RawMessage) {
	x := *d
	//removing bracket
	x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
	//removing tabs
	x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9}, []byte{})
	//removing starting quotes
	x = bytes.ReplaceAll(x, []byte{10, 34}, []byte{10})
	//removing slashes (\) for escaping quotes
	x = bytes.ReplaceAll(x, []byte{92, 34}, []byte{34})
	//removing end quote and comma
	x = bytes.ReplaceAll(x, []byte{34, 44, 10}, []byte{10})
	//removing trailing tabs
	x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 93}, []byte{})
	fmt.Printf("%s\n", x)
}

/*
I need to test the perf of followng
1. read file then unmarshal whole struct vs unmarshal 1 struct
a=> One ReadFile = size of the file, unmarshaling is negligible
2. multiple small unmarshal and sharing readfile vs 1 big unmarshal and sharing objects
*/
