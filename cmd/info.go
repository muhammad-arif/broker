/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mirantis/powerplug/misc"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
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
		case "clusterinfo":
			GetInfoCluster(&nodeDsinfoStruct)
		case "dmidecode":
			GetInfoDmidecode(&nodeDsinfoStruct)
		case "images":
			GetInfoImages(&nodeDsinfoStruct)
		case "dockerinfo":
			GetInfoDocker(&nodeDsinfoStruct)
		case "pidlimits":
			GetInfoPidLimits(&nodeDsinfoStruct)
		case "version":
			GetInfoVersion(&nodeDsinfoStruct)
		case "filemax":
			GetInfoFileMax(&nodeDsinfoStruct)
		case "ifconfig":
			GetInfoIfconfig(&nodeDsinfoStruct)
		case "kernelversion":
			GetInfoKernelVersion(&nodeDsinfoStruct)
		case "mounts":
			GetInfoMount(&nodeDsinfoStruct)
		case "networking":
			GetInfoNetworking(&nodeDsinfoStruct)
		case "psauxdocker":
			GetInfoPsAuxDocker(&nodeDsinfoStruct)
		case "rethinkstatus":
			GetInfoRethinkStatus(&nodeDsinfoStruct)
		case "running-cgroup":
			GetInfoRunningCgroup(&nodeDsinfoStruct)
		case "sestatus":
			GetInfoSestatus(&nodeDsinfoStruct)
		case "kernelconfig":
			GetInfoKernelConf(&nodeDsinfoStruct)
		case "ssd":
			GetInfoSSD(&nodeDsinfoStruct)
		case "system-cgroup":
			GetInfoSystemCgroups(&nodeDsinfoStruct)
		case "system-release":
			GetInfoSystemRelease(&nodeDsinfoStruct)
		case "ucp-controller-diag":
			GetInfoUCPControllerDiag(&nodeDsinfoStruct)
		case "uptime":
			GetInfoUptime(&nodeDsinfoStruct)
		case "daemondotjson":
			GetConfDaemonDotJson(&nodeDsinfoStruct)

		}
	}
}
func GetConfDaemonDotJson(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoCluster(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoDmidecode(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoImages(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoDocker(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoPidLimits(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoVersion(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoFileMax(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoIfconfig(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoKernelConf(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoKernelVersion(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoMount(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoNetworking(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoPsAuxDocker(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoRethinkStatus(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoRunningCgroup(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoSestatus(d *dLib.DsinfoSlashDsinfoDotJson) {

}
func GetInfoSSD(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoSystemCgroups(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoSystemRelease(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoUCPControllerDiag(d *dLib.DsinfoSlashDsinfoDotJson) {

}

func GetInfoUptime(d *dLib.DsinfoSlashDsinfoDotJson) {

}
