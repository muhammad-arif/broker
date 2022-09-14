/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
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
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"
)

var validLogArgs = map[string]int{"calico": 4, "deadcontainermounts": 4, "dmesg": 4, "dtr": 4, "daemon": 4, "kernel": 4, "shimlogs": 4}
var validLogArgsSlice = []string{"calico", "deadcontainermounts", "dmesg", "dtr", "daemon", "kernel", "shimlogs"}

// nodeCmd represents the node command.
var nodelogCmd = &cobra.Command{
	Use:   "node [arguments] [nodeName] " + "\n\nAllowed arguments are: \n\t" + strings.Join(validLogArgsSlice, "\n\t"),
	Short: "Fetch the node specific logs",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a valid argument")
		} else if len(args) < 2 {

		} else if _, ok := validLogArgs[args[0]]; ok {
			return nil
		}
		return fmt.Errorf("requires a valid argument %v", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		getLogs(args)
	},
}

func init() {
}

func getLogs(a []string) {
	node := a[1]
	log := a[0]
	nodeList, _, dsinfoJson := misc.ParseUcpNodesInspect()
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
		switch log {
		case "calico":
			GetNodeCalico(&nodeDsinfoStruct)
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
	type nestedDsinfoJournalDaemon struct {
		Calico struct {
			BgpPeers         json.RawMessage `json:"bgp_peers"`
			NodeStatus       json.RawMessage `json:"node_status"`
			WorkloadEndpoint json.RawMessage `json:"workload_endpoint"`
		} `json:"calico"`
	}
	var NodeContents nestedDsinfoJournalDaemon
	err := sonic.Unmarshal(d.DsinfoContents, &NodeContents)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	var wg = &sync.WaitGroup{}
	//fmt.Printf("%s", fmt.Sprintf("%s", NodeContents.Calico))
	wg.Add(3)
	go func() {
		x := NodeContents.Calico.BgpPeers
		//	removings /n
		x = bytes.ReplaceAll(x, []byte{92, 34}, []byte{34})
		x = bytes.ReplaceAll(x, []byte{92, 110}, []byte{10})
		x = bytes.TrimLeft(bytes.TrimRight(x, `"`), `"`)
		x = bytes.TrimFunc(x, unicode.IsMark)
		fmt.Println(string(x))
		wg.Done()
	}()
	go func() {
		x := NodeContents.Calico.NodeStatus
		//removing bracket
		x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
		//removing tabs
		x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9}, []byte{})
		x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 9, 9, 93}, []byte{})
		x = bytes.ReplaceAll(x, []byte{34}, []byte{})
		fmt.Println(string(NodeContents.Calico.NodeStatus))
		wg.Done()
	}()
	go func() {
		x := NodeContents.Calico.WorkloadEndpoint
		x = bytes.ReplaceAll(x, []byte{92, 34}, []byte{34})
		x = bytes.ReplaceAll(x, []byte{92, 110}, []byte{10})
		x = bytes.TrimLeft(bytes.TrimRight(x, `"`), `"`)
		x = bytes.TrimFunc(x, unicode.IsMark)
		fmt.Println(string(x))
		wg.Done()
	}()
	wg.Wait()
}

func GetLogsDeadContainerMounts(d *dLib.DsinfoSlashDsinfoDotJson) {
	type nestedDsinfoJournalDaemon struct {
		DeadContainerMounts json.RawMessage `json:"dead_container_mounts"`
	}
	var NodeContents nestedDsinfoJournalDaemon
	err := sonic.Unmarshal(d.DsinfoContents, &NodeContents)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	misc.PrettyPrintFromByteSlice(&NodeContents.DeadContainerMounts)

}

func GetLogsDmesg(d *dLib.DsinfoSlashDsinfoDotJson) {
	type nestedDsinfoJournalDaemon struct {
		Dmesg []struct {
			MESSAGE                  string `json:"MESSAGE"`
			PRIORITY                 string `json:"PRIORITY"`
			SYSLOGFACILITY           string `json:"SYSLOG_FACILITY"`
			SYSLOGIDENTIFIER         string `json:"SYSLOG_IDENTIFIER"`
			BOOTID                   string `json:"_BOOT_ID"`
			HOSTNAME                 string `json:"_HOSTNAME"`
			MACHINEID                string `json:"_MACHINE_ID"`
			SOURCEMONOTONICTIMESTAMP string `json:"_SOURCE_MONOTONIC_TIMESTAMP"`
			TRANSPORT                string `json:"_TRANSPORT"`
			CURSOR                   string `json:"__CURSOR"`
			MONOTONICTIMESTAMP       string `json:"__MONOTONIC_TIMESTAMP"`
			REALTIMETIMESTAMP        string `json:"__REALTIME_TIMESTAMP"`
		} `json:"dmesg"`
	}
	var NodeContents nestedDsinfoJournalDaemon
	err := sonic.Unmarshal(d.DsinfoContents, &NodeContents)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	for i, _ := range NodeContents.Dmesg {
		i64, _ := strconv.ParseInt(NodeContents.Dmesg[i].REALTIMETIMESTAMP, 10, 32)
		unixTimeUTC := time.Unix(i64, 0) //gives unix time stamp in utc
		fmt.Println(unixTimeUTC.Format(time.RFC3339), " ", NodeContents.Dmesg[i].PRIORITY, " ", NodeContents.Dmesg[i].TRANSPORT, "  ", NodeContents.Dmesg[i].HOSTNAME, "  ", NodeContents.Dmesg[i].MESSAGE)
	}
}

func GetLogsDtr(d *dLib.DsinfoSlashDsinfoDotJson) {
	type nestedDsinfoJournalDaemon struct {
		Dtr json.RawMessage `json:"dtr"`
	}
	var NodeContents nestedDsinfoJournalDaemon
	err := sonic.Unmarshal(d.DsinfoContents, &NodeContents)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Println(string(NodeContents.Dtr))
}

func GetLogsDaemon(d *dLib.DsinfoSlashDsinfoDotJson) {
	type nestedDsinfoJournalDaemon struct {
		JournalctlDaemon json.RawMessage `json:"journalctl_daemon"`
	}
	var NodeContents nestedDsinfoJournalDaemon
	err := sonic.Unmarshal(d.DsinfoContents, &NodeContents)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	misc.PrettyPrintFromByteSlice(&NodeContents.JournalctlDaemon)
}

func GetLogsKernel(d *dLib.DsinfoSlashDsinfoDotJson) {
	type nestedDsinfoJournalDaemon struct {
		JournalctlKernelDmesg []struct {
			MESSAGE                  string `json:"MESSAGE"`
			PRIORITY                 string `json:"PRIORITY"`
			SYSLOGFACILITY           string `json:"SYSLOG_FACILITY"`
			SYSLOGIDENTIFIER         string `json:"SYSLOG_IDENTIFIER"`
			BOOTID                   string `json:"_BOOT_ID"`
			HOSTNAME                 string `json:"_HOSTNAME"`
			MACHINEID                string `json:"_MACHINE_ID"`
			SOURCEMONOTONICTIMESTAMP string `json:"_SOURCE_MONOTONIC_TIMESTAMP"`
			TRANSPORT                string `json:"_TRANSPORT"`
			CURSOR                   string `json:"__CURSOR"`
			MONOTONICTIMESTAMP       string `json:"__MONOTONIC_TIMESTAMP"`
			REALTIMETIMESTAMP        string `json:"__REALTIME_TIMESTAMP"`
		} `json:"journalctl_kernel"`
	}
	var NodeContents nestedDsinfoJournalDaemon
	err := sonic.Unmarshal(d.DsinfoContents, &NodeContents)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	for i, _ := range NodeContents.JournalctlKernelDmesg {
		i64, _ := strconv.ParseInt(NodeContents.JournalctlKernelDmesg[i].REALTIMETIMESTAMP, 10, 32)
		unixTimeUTC := time.Unix(i64, 0) //gives unix time stamp in utc
		fmt.Println(unixTimeUTC.Format(time.RFC3339), " ", NodeContents.JournalctlKernelDmesg[i].PRIORITY, " ", NodeContents.JournalctlKernelDmesg[i].TRANSPORT, "  ", NodeContents.JournalctlKernelDmesg[i].HOSTNAME, "  ", NodeContents.JournalctlKernelDmesg[i].MESSAGE)
	}

}
func GetLogsShimLogs(d *dLib.DsinfoSlashDsinfoDotJson) {
	type nestedDsinfoJournalDaemon struct {
		ShimLogs map[string]json.RawMessage `json:"shim_logs"`
		//ShimLogs json.RawMessage `json:"shim_logs"`
	}
	var NodeContents nestedDsinfoJournalDaemon
	err := sonic.Unmarshal(d.DsinfoContents, &NodeContents)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	for k, _ := range NodeContents.ShimLogs {
		fmt.Println("\n-----------------------------------------------------\n", k, "\n-----------------------------------------------------\n")
		x := NodeContents.ShimLogs[k]
		//removing bracket
		x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
		x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9}, []byte{})
		x = bytes.ReplaceAll(x, []byte{34}, []byte{})
		x = bytes.ReplaceAll(x, []byte{92, 110, 44, 10}, []byte{10})

		fmt.Println(string(x))
	}
}
