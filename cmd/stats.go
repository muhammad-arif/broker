/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/mirantis/broker/misc"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"strings"

	"github.com/spf13/cobra"
)

var validStatsArgs = map[string]int{"df": 1, "dockerstats": 0, "iostat": 1, "meminfo": 2, "netstat": 3, "vmstat": 1}
var validStatsArgsSlice = []string{"df", "dockerstats", "iostat", "meminfo", "netstat", "vmstat"}

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats [argument] [nodeName]" + "\n\nAllowed arguments are: \n\t" + strings.Join(validStatsArgsSlice, "\n\t"),
	Short: "Shows stats about various node specific components",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a valid argument")
		} else if _, ok := validStatsArgs[args[0]]; ok && len(args) <= 2 {
			return nil
		}
		return fmt.Errorf("requires a valid argument")
	},

	Run: func(cmd *cobra.Command, args []string) {
		getStats(args)
	},
}

func getStats(a []string) {
	arg := a[0]
	nodeList, _, dsinfoJson := misc.ParseUcpNodesInspect()
	if len(a) == 1 {
		// cannot make the for loop concurrent because of the heavy memory usage
		//var wg = &sync.WaitGroup{}
		for k, _ := range nodeList {
			//wg.Add(1)
			//go func() {
			fmt.Printf("\n\nNode : %v\n", k)
			var nestedDsinfo = make(map[string]json.RawMessage)
			err := sonic.Unmarshal(*dsinfoJson, &nestedDsinfo)
			if err != nil {
				fmt.Errorf("cannot unmarshal %s", err)
			}
			var nodeDsinfoStruct dLib.DsinfoSlashDsinfoDotJson
			err = sonic.Unmarshal(nestedDsinfo[k], &nodeDsinfoStruct)
			if err != nil {
				fmt.Errorf("Cannot unmarshal nesteddsinfo")
			}
			switch arg {
			case "df":
				GetStatsDf(&nodeDsinfoStruct.DsinfoContents)
			case "dockerstats":
				GetStatsDockerStats(&nodeDsinfoStruct.DsinfoContents)
			case "iostat":
				GetStatsIoStat(&nodeDsinfoStruct.DsinfoContents)
			case "meminfo":
				GetStatsMeminfo(&nodeDsinfoStruct.DsinfoContents)
			case "netstat":
				GetStatsNetstats(&nodeDsinfoStruct.DsinfoContents)
			case "vmstat":
				GetStatsVmstat(&nodeDsinfoStruct.DsinfoContents)
			}
		}
	} else if len(a) == 2 {
		node := a[1]
		if _, ok := nodeList[node]; ok {
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
			switch arg {
			case "df":
				GetStatsDf(&nodeDsinfoStruct.DsinfoContents)
			case "dockerstats":
				GetStatsDockerStats(&nodeDsinfoStruct.DsinfoContents)
			case "iostat":
				GetStatsIoStat(&nodeDsinfoStruct.DsinfoContents)
			case "meminfo":
				GetStatsMeminfo(&nodeDsinfoStruct.DsinfoContents)
			case "netstat":
				GetStatsNetstats(&nodeDsinfoStruct.DsinfoContents)
			case "vmstat":
				GetStatsVmstat(&nodeDsinfoStruct.DsinfoContents)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
func GetStatsDf(d *json.RawMessage) {
	type nestedDsinfoForDf struct {
		Df json.RawMessage `json:"df"`
	}
	var NodeContents nestedDsinfoForDf
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	misc.PrettyPrintFromByteSlice(&NodeContents.Df)
}

func GetStatsDockerStats(d *json.RawMessage) {
	type stats struct {
		BlockIO   string `json:"BlockIO"`
		CPUPerc   string `json:"CPUPerc"`
		Container string `json:"Container"`
		ID        string `json:"ID"`
		MemPerc   string `json:"MemPerc"`
		MemUsage  string `json:"MemUsage"`
		Name      string `json:"Name"`
		NetIO     string `json:"NetIO"`
		PIDs      string `json:"PIDs"`
	}
	type nestedDsinfoForDf struct {
		DockerStats []stats `json:"docker_stats"`
	}
	var NodeContents nestedDsinfoForDf
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AlignHeader: text.AlignCenter},
		{Number: 2, AlignHeader: text.AlignCenter},
		{Number: 3, AlignHeader: text.AlignCenter},
		{Number: 4, AlignHeader: text.AlignCenter},
		{Number: 5, AlignHeader: text.AlignCenter},
		{Number: 6, AlignHeader: text.AlignCenter},
	})
	t.AppendHeader(table.Row{"Name", "PID", "CPUPerc", "MemPerc", "MemUsage", "BlockIO", "NetIO"})

	for i, _ := range NodeContents.DockerStats {
		t.AppendRow(table.Row{fmt.Sprintf("%.35s", NodeContents.DockerStats[i].Name), NodeContents.DockerStats[i].PIDs, NodeContents.DockerStats[i].CPUPerc, NodeContents.DockerStats[i].MemPerc, NodeContents.DockerStats[i].MemUsage, NodeContents.DockerStats[i].BlockIO, NodeContents.DockerStats[i].NetIO})
	}
	fmt.Println(t.Render())

}

func GetStatsIoStat(d *json.RawMessage) {
	type nestedDsinfoForDf struct {
		Iostat json.RawMessage `json:"iostat"`
	}
	var NodeContents nestedDsinfoForDf
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	misc.PrettyPrintFromByteSlice(&NodeContents.Iostat)
}

func GetStatsMeminfo(d *json.RawMessage) {
	type nestedDsinfoForDf struct {
		Meminfo json.RawMessage `json:"meminfo"`
	}
	var NodeContents nestedDsinfoForDf
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	misc.PrettyPrintFromByteSlice(&NodeContents.Meminfo)
}

func GetStatsNetstats(d *json.RawMessage) {
	type nestedDsinfoForDf struct {
		Netstat json.RawMessage `json:"netstat"`
	}
	var NodeContents nestedDsinfoForDf
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	misc.PrettyPrintFromByteSlice(&NodeContents.Netstat)
}
func GetStatsVmstat(d *json.RawMessage) {
	type nestedDsinfoForDf struct {
		Vmstat json.RawMessage `json:"vmstat"`
	}
	var NodeContents nestedDsinfoForDf
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	misc.PrettyPrintFromByteSlice(&NodeContents.Vmstat)
}
