/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/
package container

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mirantis/broker/misc"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

var verbose bool

// listCmd represents the list command
var ContListCmd = &cobra.Command{
	Use:   "ls [all|NODE NAME]",
	Short: "List containers",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			GetContainers([]string{"all"})
			return
		} else {
			GetContainers(args)
		}
	},
	Aliases: []string{"list", "ls", "l", "ps"},
}

func init() {
	ContListCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}

//minimal type for efficient unmarshaling
type NestedDsinfoT struct {
	ContainerInfo map[string]containerInfoT `json:"container_info"`
}
type containerInfoT struct {
	Inspect types.ContainerJSON `json:"inspect"`
}

func GetContainers(a []string) {
	nodeList, _, dsinfoJson := misc.ParseUcpNodesInspect()
	if a[0] == "all" {
		var nestedDsinfo = make(map[string]json.RawMessage)
		err := sonic.Unmarshal(*dsinfoJson, &nestedDsinfo)
		if err != nil {
			fmt.Errorf("cannot unmarshal %s", err)
		}
		//Ignoring concurrency as it seems 3s slower than the usual
		//var wg = &sync.WaitGroup{}
		for k, _ := range nodeList {
			//wg.Add(1)
			//go func() {
			createPerNodeContainerList(k, nestedDsinfo)
			//wg.Done()
			//}()
		}
		//wg.Wait()
	} else {
		for _, k := range a {
			//check arguments validity
			if _, ok := nodeList[k]; ok {
				var nestedDsinfo = make(map[string]json.RawMessage)
				err := sonic.Unmarshal(*dsinfoJson, &nestedDsinfo)
				if err != nil {
					fmt.Errorf("cannot unmarshal %s", err)
				}
				createPerNodeContainerList(k, nestedDsinfo)
			}
		}
	}
}
func createPerNodeContainerList(k string, nestedDsinfo map[string]json.RawMessage) {
	t := table.NewWriter()
	//t.SetAllowedRowLength(200)
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AlignHeader: text.AlignCenter},
		{Number: 2, AlignHeader: text.AlignCenter},
		{Number: 3, AlignHeader: text.AlignCenter},
		{Number: 4, AlignHeader: text.AlignCenter},
		{Number: 5, AlignHeader: text.AlignCenter},
		{Number: 6, AlignHeader: text.AlignCenter},
		{Number: 7, AlignHeader: text.AlignCenter},
	})
	t.SetTitle(k)
	t.AppendHeader(table.Row{"Name", "Image", "Created", "Status", "OOMK", "Rst-ng", "Network"})
	var nodeDsinfoStruct dLib.DsinfoSlashDsinfoDotJson
	err := sonic.Unmarshal(nestedDsinfo[k], &nodeDsinfoStruct)
	if err != nil {
		fmt.Errorf("Cannot unmarshal nesteddsinfo")
	}
	var NodeContents NestedDsinfoT
	err = sonic.Unmarshal(nodeDsinfoStruct.DsinfoContents, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	for _, v := range NodeContents.ContainerInfo {
		cName, iName, nName := dressingRoom(v.Inspect.Name, v.Inspect.Config.Image, v.Inspect.NetworkSettings.Networks)
		t.AppendRow(table.Row{cName, iName, fmt.Sprintf("%.19s", v.Inspect.Created), v.Inspect.State.Status, v.Inspect.State.Restarting, v.Inspect.State.OOMKilled, fmt.Sprintf("%.20s", nName)})

	}
	fmt.Println(t.Render())

}
func dressingRoom(c string, i string, n map[string]*network.EndpointSettings) (string, string, string) {
	// [nName] Just collecting the keys
	modifiedN := []string{}
	for k, _ := range n {
		modifiedN = append(modifiedN, k)
	}
	// converting the modifiedN to a comma seperated string
	css := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(modifiedN)), ", "), "[]")

	// [cName] Removing openning slash `/`
	re := regexp.MustCompile(`/(.*$)`)
	cName := re.FindSubmatch([]byte(c))
	// if flag is set to verbose=true then we keep the original size
	if verbose {
		return string(cName[1]), i, css
	}
	// [cName] Reducing no. of characters
	modifiedC := fmt.Sprintf("%.35s", cName[1])

	// [iName] removing registry and digest
	var modifiedI string
	if re := regexp.MustCompile(`sha...`); re.MatchString(i) {
		if re := regexp.MustCompile(`^sha...:.*$`); re.MatchString(i) {
			modifiedI = fmt.Sprintf("%.15s", i)
		} else if re := regexp.MustCompile(`.*@sha...:[a-z0-9]+$`); re.MatchString(i) {
			xe := regexp.MustCompile(`/([a-z0-9-]+@sha...:.....)`)
			x := xe.FindSubmatch([]byte(i))
			if len(x) == 2 {
				modifiedI = string(x[1])
			} else {
				modifiedI = ""
			}
		} else {
			modifiedI = ""
		}
	} else if re := regexp.MustCompile(`/[a-z0-9-]+:[a-z0-9-.]+`); re.MatchString(i) {
		xe := regexp.MustCompile(`/([a-z0-9-]+:[a-z0-9-.]+)`)
		x := xe.FindSubmatch([]byte(i))
		if len(x) == 2 {
			modifiedI = string(x[1])
		} else {
			modifiedI = ""
		}

	}

	return modifiedC, modifiedI, css
}
