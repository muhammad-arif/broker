/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/docker/docker/api/types"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/mirantis/broker/misc"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"github.com/spf13/cobra"
	"strings"
)

var pretty bool
var verbose bool

var validInfoArgs = map[string]int{"lscpu": 0, "clusterinfo": 1, "dmidecode": 2, "images": 3, "dockerinfo": 4, "pidlimits": 5, "version": 6, "filemax": 7, "ifconfig": 8, "kernel": 9, "mounts": 11, "networking": 12, "psauxdocker": 12, "rethinkstatus": 13, "running-cgroup": 14, "sestatus": 15, "kernelconf": 16, "ssd": 17, "system-cgroup": 18, "system-release": 19, "ucp-controller-diag": 20, "uptime": 21, "daemondotjson": 22}
var validInfoArgsSlice = []string{"lscpu", "clusterinfo", "dmidecode", "images", "dockerinfo", "pidlimits", "version", "filemax", "ifconfig", "kernel", "mounts", "networking", "psauxdocker", "rethinkstatus", "running-cgroup", "sestatus", "kernelconf", "ssd", "system-cgroup", "system-release", "ucp-controller-diag", "uptime", "daemondotjson"}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use: "info [argument] [nodeName]" + "\n\nAllowed arguments are: \n\t" + strings.Join(validInfoArgsSlice, "\n\t"),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires a valid argument")
		} else if _, ok := validInfoArgs[args[0]]; ok && len(args) <= 2 {
			return nil
		}
		return fmt.Errorf("requires a valid argument")
	},
	Short: "Shows information about various node specific components",
	Run: func(cmd *cobra.Command, args []string) {
		getInfo(args)
	},
	Aliases: []string{"info", "i", "inf"},
}

func init() {
	infoCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "pretty JSON output")
	infoCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

}

func getInfo(a []string) {
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
			switch a[0] {
			case "clusterinfo":
				GetInfoCluster(&nodeDsinfoStruct.DsinfoContents)
			case "dmidecode":
				GetInfoDmidecode(&nodeDsinfoStruct.DsinfoContents)
			case "images":
				GetInfoImages(&nodeDsinfoStruct.DsinfoContents)
			case "dockerinfo":
				GetInfoDocker(&nodeDsinfoStruct.DsinfoContents)
			case "pidlimits":
				GetInfoPidLimits(&nodeDsinfoStruct.DsinfoContents)
			case "version":
				GetInfoVersion(&nodeDsinfoStruct.DsinfoContents)
			case "filemax":
				GetInfoFileMax(&nodeDsinfoStruct.DsinfoContents)
			case "ifconfig":
				GetInfoIfconfig(&nodeDsinfoStruct.DsinfoContents)
			case "kernel":
				GetInfoKernelVersion(&nodeDsinfoStruct.DsinfoContents)
			case "mounts":
				GetInfoMount(&nodeDsinfoStruct.DsinfoContents)
			case "psauxdocker":
				GetInfoPsAuxDocker(&nodeDsinfoStruct.DsinfoContents)
			case "rethinkstatus":
				GetInfoRethinkStatus(&nodeDsinfoStruct.DsinfoContents)
			case "running-cgroup":
				GetInfoRunningCgroup(&nodeDsinfoStruct.DsinfoContents)
			case "sestatus":
				GetInfoSestatus(&nodeDsinfoStruct.DsinfoContents)
			case "kernelconf":
				GetInfoKernelConf(&nodeDsinfoStruct.DsinfoContents)
			case "ssd":
				GetInfoSSD(&nodeDsinfoStruct.DsinfoContents)
			case "system-cgroup":
				GetInfoSystemCgroups(&nodeDsinfoStruct.DsinfoContents)
			case "system-release":
				GetInfoSystemRelease(&nodeDsinfoStruct.DsinfoContents)
			case "ucp-controller-diag":
				GetInfoUCPControllerDiag(&nodeDsinfoStruct.DsinfoContents)
			case "uptime":
				GetInfoUptime(&nodeDsinfoStruct.DsinfoContents)
			case "daemondotjson":
				GetConfDaemonDotJson(&nodeDsinfoStruct.DsinfoContents)
			case "lscpu":
				GetInfoLscpu(&nodeDsinfoStruct.DsinfoContents)
			}
			//wg.Done()
			//}()

		}
		//wg.Wait()
	} else if len(a) == 2 {
		arg := a[0]
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
			case "clusterinfo":
				GetInfoCluster(&nodeDsinfoStruct.DsinfoContents)
			case "dmidecode":
				GetInfoDmidecode(&nodeDsinfoStruct.DsinfoContents)
			case "images":
				GetInfoImages(&nodeDsinfoStruct.DsinfoContents)
			case "dockerinfo":
				GetInfoDocker(&nodeDsinfoStruct.DsinfoContents)
				//fmt.Println("Invoked GetInfoDocker")
			case "pidlimits":
				GetInfoPidLimits(&nodeDsinfoStruct.DsinfoContents)
			case "version":
				GetInfoVersion(&nodeDsinfoStruct.DsinfoContents)
			case "filemax":
				GetInfoFileMax(&nodeDsinfoStruct.DsinfoContents)
			case "ifconfig":
				GetInfoIfconfig(&nodeDsinfoStruct.DsinfoContents)
			case "kernel":
				GetInfoKernelVersion(&nodeDsinfoStruct.DsinfoContents)
			case "mounts":
				GetInfoMount(&nodeDsinfoStruct.DsinfoContents)
			case "psauxdocker":
				GetInfoPsAuxDocker(&nodeDsinfoStruct.DsinfoContents)
			case "rethinkstatus":
				GetInfoRethinkStatus(&nodeDsinfoStruct.DsinfoContents)
			case "running-cgroup":
				GetInfoRunningCgroup(&nodeDsinfoStruct.DsinfoContents)
			case "sestatus":
				GetInfoSestatus(&nodeDsinfoStruct.DsinfoContents)
			case "kernelconf":
				GetInfoKernelConf(&nodeDsinfoStruct.DsinfoContents)
			case "ssd":
				GetInfoSSD(&nodeDsinfoStruct.DsinfoContents)
			case "system-cgroup":
				GetInfoSystemCgroups(&nodeDsinfoStruct.DsinfoContents)
			case "system-release":
				GetInfoSystemRelease(&nodeDsinfoStruct.DsinfoContents)
			case "ucp-controller-diag":
				GetInfoUCPControllerDiag(&nodeDsinfoStruct.DsinfoContents)
			case "uptime":
				GetInfoUptime(&nodeDsinfoStruct.DsinfoContents)
			case "daemondotjson":
				GetConfDaemonDotJson(&nodeDsinfoStruct.DsinfoContents)
			case "lscpu":
				GetInfoLscpu(&nodeDsinfoStruct.DsinfoContents)
			}
		}
	}

}

// GetInfoLscpu shows list of cpus
func GetInfoLscpu(d *json.RawMessage) {
	type nestedDsinfoForDf struct {
		Lscpu json.RawMessage `json:"lscpu"`
	}
	var NodeContents nestedDsinfoForDf
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	misc.PrettyPrintFromByteSlice(&NodeContents.Lscpu)

}

// GetConfDaemonDotJson Shows daemon.json
func GetConfDaemonDotJson(d *json.RawMessage) {
	type nestedDsinfoForDf struct {
		DockerDaemonJson json.RawMessage `json:"docker_daemon_json"`
	}
	var NodeContents nestedDsinfoForDf
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x, err := json.Marshal(NodeContents.DockerDaemonJson)
	if err != nil {
		fmt.Errorf("not unmarshalled properly, %v", err)
	}
	if pretty {
		z, _ := misc.PrettyString(string(x))
		fmt.Println(z)
	} else {
		fmt.Println(string(x))
	}
}

// GetInfoCluster Shows cluster info from leader
func GetInfoCluster(d *json.RawMessage) {
	type nestedDsinfoForDf struct {
		//DockerInfo json.RawMessage `json:"docker_info"`
		ClusterInfo json.RawMessage `json:"cluster_info"`
	}
	var NodeContents nestedDsinfoForDf
	//sonic.Decoder.Decode(strings(*d.DsinfoContents))
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x, err := json.Marshal(NodeContents.ClusterInfo)
	if err != nil {
		fmt.Errorf("not unmarshalled properly, %v", err)
	}
	if pretty {
		z, _ := misc.PrettyString(string(x))
		fmt.Println(z)
	} else {
		fmt.Println(string(x))
	}
}

// GetInfoDmidecode shows `dmidecode`
func GetInfoDmidecode(d *json.RawMessage) {
	type nestedDsinfoForDf struct {
		Dmidecode json.RawMessage `json:"dmidecode"`
	}
	var NodeContents nestedDsinfoForDf
	//sonic.Decoder.Decode(strings(*d.DsinfoContents))
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}

	//removing \t, quotes and comma
	fmt.Printf("%s\n", bytes.ReplaceAll(bytes.ReplaceAll(bytes.ReplaceAll(bytes.ReplaceAll(NodeContents.Dmidecode, []byte(`\t`), []byte("    ")), []byte(`	`), []byte("")), []byte(`"`), []byte(``)), []byte{44, 10}, []byte{10}))
}

// GetInfoImages shows images of a particular node
func GetInfoImages(d *json.RawMessage) {
	type nestedDsinfoForImage struct {
		DockerImages []misc.ImageMetadata `json:"docker_images"`
	}
	var NodeContents nestedDsinfoForImage
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	if pretty {
		//NodeContents.DockerImages
	} else {
		t := table.NewWriter()
		t.AppendHeader(table.Row{"Repo", "Tag", "ID", "Created At", "Size"})
		for _, v := range NodeContents.DockerImages {
			t.AppendRow(table.Row{v.Repository, v.Tag, misc.TruncateID(v.ID), v.CreatedAt, v.Size})
		}
		fmt.Println(t.Render())
	}
}

// GetInfoDocker shows docker info
func GetInfoDocker(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		DockerInfo json.RawMessage `json:"docker_info"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x, err := json.Marshal(NodeContents.DockerInfo)
	if err != nil {
		fmt.Errorf("not unmarshalled properly, %v", err)
	}
	if pretty {
		z, _ := misc.PrettyString(string(x))
		fmt.Println(z)
	} else {
		fmt.Println(string(x))
	}
}

// GetInfoKernelConf shows a few kernel configurations
func GetInfoKernelConf(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		KernelConfig json.RawMessage `json:"kernel_config"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	fmt.Println(string(NodeContents.KernelConfig))
}

// GetInfoPidLimits will return docker pid limits
func GetInfoPidLimits(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		DockerPidLimits json.RawMessage `json:"docker_pid_limits"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x := NodeContents.DockerPidLimits
	//removing bracket
	x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
	//removing tabs
	x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9}, []byte{})
	//removing quotes
	x = bytes.ReplaceAll(x, []byte{34}, []byte{})
	//removing end commas
	x = bytes.ReplaceAll(x, []byte{44, 10}, []byte{10})
	//removing trailing tabs
	x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 93}, []byte{})
	fmt.Printf("%s\n", x)
}

// GetInfoVersion shows the docker cli and server version
func GetInfoVersion(d *json.RawMessage) {

	if verbose {
		type nestedDsinfoForInfoDocker struct {
			DockerVersion json.RawMessage `json:"docker_version"`
		}
		var NodeContents nestedDsinfoForInfoDocker
		err := sonic.Unmarshal(*d, &NodeContents)
		if err != nil {
			fmt.Errorf("Cannot unmarshal NodeContents")
		}
		fmt.Println(string(NodeContents.DockerVersion))
	} else {
		type nestedDsinfoForInfoDocker struct {
			DockerVersion map[string]types.Version `json:"docker_version"`
		}
		var NodeContents nestedDsinfoForInfoDocker
		err := sonic.Unmarshal(*d, &NodeContents)
		if err != nil {
			fmt.Errorf("Cannot unmarshal NodeContents")
		}
		fmt.Printf("Client: \n\tVersion: %v \n\tAPI Version: %v\n", NodeContents.DockerVersion["Client"].Version, NodeContents.DockerVersion["Client"].APIVersion)
		fmt.Printf("Server: \n\tVersion: %v \n\tAPI Version: %v\n\tLicense : %v\n", NodeContents.DockerVersion["Server"].Version, NodeContents.DockerVersion["Server"].APIVersion, NodeContents.DockerVersion["Server"].Platform.Name)
		fmt.Printf("Kernel: \n\tVersion: %v \n", NodeContents.DockerVersion["Server"].KernelVersion)

	}

}

// GetInfoFileMax will return filemax
func GetInfoFileMax(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		FileMax json.RawMessage `json:"file_max"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	fmt.Println(string(NodeContents.FileMax))
}

// GetInfoIfconfig shows if config of a host
func GetInfoIfconfig(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		Ifconfig json.RawMessage `json:"ifconfig"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	fmt.Println(string(NodeContents.Ifconfig))
}

// GetInfoKernelVersion
func GetInfoKernelVersion(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		KernelVersion json.RawMessage `json:"kernel_version"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x := NodeContents.KernelVersion
	//removing bracket
	x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
	//removing tabs
	x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9}, []byte{})
	//removing starting quotes
	x = bytes.ReplaceAll(x, []byte{10, 34}, []byte{10})
	//removing end quote
	x = bytes.ReplaceAll(x, []byte{34, 10}, []byte{10})
	//removing trailing tabs
	x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 93}, []byte{})
	fmt.Printf("%s\n", x)
}

// GetInfoMount
func GetInfoMount(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		Mount json.RawMessage `json:"mount"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	misc.PrettyPrintFromByteSlice(&NodeContents.Mount)
}

func GetInfoPsAuxDocker(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		PsAuxGrepDocker json.RawMessage `json:"ps_aux_grep_docker"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	fmt.Println(string(NodeContents.PsAuxGrepDocker))
	//misc.PrettyPrintFromByteSlice(&NodeContents.PsAuxGrepDocker)

}

func GetInfoRethinkStatus(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		RethinkStatus json.RawMessage `json:"rethink_status"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	fmt.Println(string(NodeContents.RethinkStatus))

}

func GetInfoRunningCgroup(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		RunningCgroup json.RawMessage `json:"running_cgroup"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x := NodeContents.RunningCgroup
	x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
	x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9}, []byte{})
	x = bytes.ReplaceAll(x, []byte{10, 34}, []byte{10})
	x = bytes.ReplaceAll(x, []byte{34, 44, 10}, []byte{10})
	x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 93}, []byte{})
	fmt.Println(string(x))
}

func GetInfoSestatus(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		Sestatus json.RawMessage `json:"sestatus"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	fmt.Println(string(NodeContents.Sestatus))

}
func GetInfoSSD(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		Ssd json.RawMessage `json:"ssd"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	fmt.Println(string(NodeContents.Ssd))

}

func GetInfoSystemCgroups(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		SystemCgroups json.RawMessage `json:"system_cgroups"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x := NodeContents.SystemCgroups
	x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
	x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9}, []byte{})
	x = bytes.ReplaceAll(x, []byte{92, 116}, []byte{9, 9})
	x = bytes.ReplaceAll(x, []byte{10, 34}, []byte{10})
	x = bytes.ReplaceAll(x, []byte{34, 44, 10}, []byte{10})
	x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 93}, []byte{})
	fmt.Println(string(x))
}

func GetInfoSystemRelease(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		SystemRelease json.RawMessage `json:"system_release"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x := NodeContents.SystemRelease
	x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
	x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9}, []byte{})
	x = bytes.ReplaceAll(x, []byte{92, 116}, []byte{9, 9})
	x = bytes.ReplaceAll(x, []byte{92, 34}, []byte{})
	x = bytes.ReplaceAll(x, []byte{10, 34}, []byte{10})
	x = bytes.ReplaceAll(x, []byte{34, 44, 10}, []byte{10})
	x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 93}, []byte{})

	fmt.Println(string(x))
}

func GetInfoUCPControllerDiag(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		UcpControllerDiag map[string]json.RawMessage `json:"ucp_controller_diag"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	for k, _ := range NodeContents.UcpControllerDiag {
		fmt.Println(k)
		x := NodeContents.UcpControllerDiag[k]
		x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
		x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9, 34}, []byte{9})
		x = bytes.ReplaceAll(x, []byte{92, 110, 34, 44, 10}, []byte{10})
		x = bytes.ReplaceAll(x, []byte{92, 116}, []byte{9})
		x = bytes.ReplaceAll(x, []byte{92, 110, 34, 10, 9, 9, 9, 9, 93}, []byte{10})
		fmt.Println(string(x))
	}
}

func GetInfoUptime(d *json.RawMessage) {
	type nestedDsinfoForInfoDocker struct {
		Vmstat json.RawMessage `json:"vmstat"`
	}
	var NodeContents nestedDsinfoForInfoDocker
	err := sonic.Unmarshal(*d, &NodeContents)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}

}
