/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package node

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	misc "github.com/mirantis/powerplug/misc"
	"github.com/spf13/cobra"
)

// nodeListCmd represents the list command
var nodeListCmd = &cobra.Command{
	Use:   "powerplug node ls",
	Short: "List swarm nodes",
	Run: func(cmd *cobra.Command, args []string) {
		NodeList()
	},
	Aliases: []string{"ls", "l", "list"},
}

func init() {
}
func NodeList() {
	_, inspect, _ := misc.ParseUcpNodesInspect()
	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, AlignHeader: text.AlignCenter},
		{Number: 2, AlignHeader: text.AlignCenter},
		{Number: 3, AlignHeader: text.AlignCenter},
		{Number: 4, AlignHeader: text.AlignCenter},
		{Number: 5, AlignHeader: text.AlignCenter},
		{Number: 6, AlignHeader: text.AlignCenter},
		{Number: 7, AlignHeader: text.AlignCenter},
		{Number: 8, AlignHeader: text.AlignCenter},
		{Number: 9, AlignHeader: text.AlignCenter},
	})
	t.AppendHeader(table.Row{"Name", "Role", "Avail", "State", "IP", "OS", "Engine", "MKE", "Orca"})
	for _, v := range *inspect {
		// Collecting info from labels
		v.CreatedAt.Date()
		orca, mkeVersion := OrchaInfoFromLabel(v.Spec.Labels)
		if v.Spec.Role == "manager" {
			if v.ManagerStatus.Leader {
				//dd, ll, kk := v.CreatedAt.Date()
				t.AppendRow(table.Row{v.Description.Hostname, "LEADER", v.Spec.Availability, v.Status.State, v.Status.Addr, v.Description.Platform.OS, v.Description.Engine.EngineVersion, mkeVersion, orca})
			} else {
				t.AppendRow(table.Row{v.Description.Hostname, v.Spec.Role, v.Spec.Availability, v.Status.State, v.Status.Addr, v.Description.Platform.OS, v.Description.Engine.EngineVersion, mkeVersion, orca})
			}
		} else {
			t.AppendRow(table.Row{v.Description.Hostname, v.Spec.Role, v.Spec.Availability, v.Status.State, v.Status.Addr, v.Description.Platform.OS, v.Description.Engine.EngineVersion, mkeVersion, orca})
		}

	}
	fmt.Println(t.Render())
}
func OrchaInfoFromLabel(l map[string]string) (string, string) {
	var orchastrator, mkeversion string = "", ""
	//if okv1, okv2 := l["com.docker.ucp.orchestrator.kubernetes"], l["com.docker.ucp.orchestrator.swarm"]; okv1 != "" && okv2 != "" {
	if l["com.docker.ucp.orchestrator.kubernetes"] == "true" && l["com.docker.ucp.orchestrator.swarm"] == "true" {
		orchastrator = "mixed"
	} else if l["com.docker.ucp.orchestrator.kubernetes"] == "true" && l["com.docker.ucp.orchestrator.swarm"] == "false" {
		orchastrator = "k8s"
	} else if l["com.docker.ucp.orchestrator.kubernetes"] == "false" && l["com.docker.ucp.orchestrator.swarm"] == "true" {
		orchastrator = "swarm"
	} else if l["com.docker.ucp.orchestrator.kubernetes"] == "" && l["com.docker.ucp.orchestrator.swarm"] == "true" {
		orchastrator = "swarm"
	} else if l["com.docker.ucp.orchestrator.kubernetes"] == "true" && l["com.docker.ucp.orchestrator.swarm"] == "" {
		orchastrator = "k8s"
	} else {
		orchastrator = ""
	}
	//}
	if val, ok := l["com.docker.ucp.node-state-augmented.reconciler-ucp-version"]; ok {
		mkeversion = val
	}
	return orchastrator, mkeversion
}
