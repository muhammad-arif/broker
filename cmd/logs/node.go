/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package logs

import (
	"encoding/json"
	"fmt"
	"github.com/mirantis/broker/misc"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"github.com/spf13/cobra"
)

// nodeCmd represents the node command.
var nodelogCmd = &cobra.Command{
	Use:   "node [OPTIONS] [nodeName] [logName]",
	Args:  cobra.MinimumNArgs(2),
	Short: "Fetch the node specific logs",
	Run: func(cmd *cobra.Command, args []string) {
		getLogs(args)
	},
}

func init() {
}

//var logs = map[string]string{
//	"calico":              "Calico",
//	"clusterinfo":         "ClusterInfo",
//	"stacktrace":          "DaemonStackTrace",
//	"deadcontainermounts": "DeadContainerMounts",
//	"df":                  "Df",
//	"dmesg":               "Dmesg",
//	"dmidecode":           "Dmidecode",
//	"daemondotjson":       "DockerDaemonJson",
//	"images":              "DockerImages",
//	"dockerinfo":          "DockerInfo",
//	"pidlimits":           "DockerPidLimits",
//	"dockerstats":         "DockerStats",
//	"version":             "DockerVersion",
//	"dtr":                 "Dtr",
//	"filemax":             "FileMax",
//	"hostname":            "Hostname",
//	"ifconfig":            "Ifconfig",
//	"iostat":              "Iostat",
//	"daemon":              "JournalctlDaemon",
//	"kernel":              "JournalctlKernel",
//	"kernelconf":          "KernelConfig",
//	"kernelversion":       "KernelVersion",
//	"lscpu":               "Lscpu",
//	"meminfo":             "Meminfo",
//	"mount":               "Mount",
//	"netstat":             "Netstat",
//	"networking":          "Networking",
//	"psauxgrepdocker":     "PsAuxGrepDocker",
//	"rethinkstatus":       "RethinkStatus",
//	"running-cgroup":      "RunningCgroup",
//	"sestatus":            "Sestatus",
//	"shimlogs":            "ShimLogs",
//	"ssd":                 "Ssd",
//	"system-cgroups":      "SystemCgroups",
//	"systemrelease":       "SystemRelease",
//	"ucpcontrollerdiag":   "UcpControllerDiag",
//	"uptime":              "Uptime",
//	"vmstat":              "Vmstat",
//}

func getLogs(a []string) {
	node := a[0]
	log := a[1]
	nodeList, _, dsinfoJson := misc.ParseUcpNodesInspect()
	if _, ok := nodeList[node]; ok {
		var nestedDsinfo = make(map[string]json.RawMessage)
		err := json.Unmarshal(*dsinfoJson, &nestedDsinfo)
		if err != nil {
			fmt.Errorf("cannot unmarshal %s", err)
		}
		var nodeDsinfoStruct dLib.DsinfoSlashDsinfoDotJson
		err = json.Unmarshal(nestedDsinfo[a[0]], &nodeDsinfoStruct)
		if err != nil {
			fmt.Errorf("Cannot unmarshal nesteddsinfo")
		}
		switch log {
		case "calico":
			GetNodeCalico(&nodeDsinfoStruct)
		case "stacktrace":
			GetLogsStackTrace(&nodeDsinfoStruct)
		case "deadcontainermounts":
			GetLogsDeadContainerMounts(&nodeDsinfoStruct)
		case "dmesg":
			GetLogsDmesg(&nodeDsinfoStruct)
		case "dtr":
			GetLogsDtr(&nodeDsinfoStruct)
		case "daemon":
			GetLogsDaemon(&nodeDsinfoStruct)
		case "kernel":
			GetLogsKernel(&nodeDsinfoStruct)
		case "shimlogs":
			GetLogsShimLogs(&nodeDsinfoStruct)
		}
	}
}

func GetNodeCalico(d *dLib.DsinfoSlashDsinfoDotJson) {
	fmt.Println("logs of calico")
}

func GetLogsStackTrace(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetLogsDeadContainerMounts(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetLogsDmesg(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetLogsDtr(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetLogsDaemon(d *dLib.DsinfoSlashDsinfoDotJson) {
	type nestedDsinfoJournalDaemon struct {
		//JournalctlDaemon json.RawMessage `json:"journalctl_daemon"`
		JournalctlDaemon []byte `json:"journalctl_daemon"`
		//JournalctlDaemon interface{} `json:"journalctl_daemon"`
	}
	type lines struct {
		l []string
	}
	var NodeContents nestedDsinfoJournalDaemon
	err := json.Unmarshal(d.DsinfoContents, &NodeContents)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	//var x []string
	//s, _ := strconv.Unquote(string(NodeContents.JournalctlDaemon))
	//s := json.Unmarshal(NodeContents.JournalctlDaemon, &x)
	//var eachline lines
	//json.Unmarshal(NodeContents.JournalctlDaemon, eachline)
	fmt.Println(NodeContents.JournalctlDaemon)

	//fmt.Println(string(NodeContents.JournalctlDaemon))
	//v:=[]byte{}
	//ls := fmt.Sprintf(string(NodeContents.JournalctlDaemon))
	//for _, v := range NodeContents.JournalctlDaemon. {
	//	fmt.Println(v)
	//xd := fmt.Sprintf("%s", NodeContents.JournalctlDaemon)
	//xd := fmt.Sprint(string(NodeContents.JournalctlDaemon))
	//fmt.Println(xd)
	//}
	//for _, v := range  {
	//	fmt.Println(string(v))
	//}

}

func GetLogsKernel(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetLogsShimLogs(d *dLib.DsinfoSlashDsinfoDotJson) {

}
