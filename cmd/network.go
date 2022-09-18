/*
Copyright © 2022 Matteo Andrii Marjan Prashant Oleksandr George Artur and all EMEA/APAC/AMER TSE Colleagues
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/mirantis/broker/misc"
	dLib "github.com/muhammad-arif/dsinfoParsingLibrary"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var validNetworkArgs = map[string]int{"ip": 0, "link": 1, "route": 2, "iptables": 3}
var validNetworkArgSlice = []string{"ip", "link", "route", "iptables"}

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network [argument] [nodeName]" + "\n\nAllowed arguments are: \n\t" + strings.Join(validNetworkArgSlice, "\n\t"),
	Short: "Shows network information of a node",
	Args: func(cmd *cobra.Command, args []string) error {
		if _, ok := validNetworkArgs[args[0]]; ok && len(args) == 2 {
			return nil
		}
		return fmt.Errorf("requires 2 valid argument")

	},
	Aliases: []string{"net", "network", "nt"},
	Run: func(cmd *cobra.Command, args []string) {
		getNetwork(args)
	},
}

func init() {
}

type nestedDsinfoForInfoDocker struct {
	Networking json.RawMessage `json:"networking"`
}

func getNetwork(a []string) {
	nodeList, _, dsinfoJson := misc.ParseUcpNodesInspect()
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
		var NodeContents nestedDsinfoForInfoDocker
		err = sonic.Unmarshal(nodeDsinfoStruct.DsinfoContents, &NodeContents)
		if err != nil {
			fmt.Errorf("Cannot unmarshal NodeContents")
		}
		//fmt.Println(string(NodeContents.Networking))
		switch arg {
		case "all":
			fmt.Println(string(NodeContents.Networking))
		case "ip":
			GetNetIP(&NodeContents.Networking)
		case "link":
			GetNetLink(&NodeContents.Networking)
		case "route":
			GetNetRoute(&NodeContents.Networking)
		case "iptables":
			GetNetIptables(&NodeContents.Networking)
		}
	} else {
		fmt.Println("wrong node name. please run `broker node ls -q` and select one")

	}
}

type dsinfoNetworkSystem struct {
	System struct {
		Ip struct {
			Ipv4      []string `json:"ipv4"`
			Ipv4Route []string `json:"ipv4_route"`
			Link      []string `json:"link"`
		} `json:"ip"`
		Iptables struct {
			Filter struct {
				Field1 []string `json:""`
			} `json:"filter"`
			Mangle struct {
				Field1 []string `json:""`
			} `json:"mangle"`
			Nat struct {
				Field1 []string `json:""`
			} `json:"nat"`
		} `json:"iptables"`
	} `json:"system"`
}

func GetNetRoute(j *json.RawMessage) {
	type route struct {
		System struct {
			Ip struct {
				Ipv4Route json.RawMessage `json:"ipv4_route"`
			} `json:"ip"`
		} `json:"system"`
	}
	var ipRoute route
	err := sonic.Unmarshal(*j, &ipRoute)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x := ipRoute.System.Ip.Ipv4Route
	//removing bracket
	x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
	//removing tabs
	x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9, 9, 9}, []byte{})
	//removing quotes
	x = bytes.ReplaceAll(x, []byte{34}, []byte{})
	//removing end commas
	x = bytes.ReplaceAll(x, []byte{44, 10}, []byte{10})
	//removing trailing tabs
	x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 9, 9, 9, 93}, []byte{})
	fmt.Printf("%s\n", x)
}

func GetNetLink(j *json.RawMessage) {
	type link struct {
		System struct {
			Ip struct {
				Link json.RawMessage `json:"link"`
			} `json:"ip"`
		} `json:"system"`
	}
	var ipLinks link
	err := sonic.Unmarshal(*j, &ipLinks)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x := ipLinks.System.Ip.Link
	//removing bracket
	x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
	//removing unicode escaped chars for <
	x = bytes.ReplaceAll(x, []byte{92, 117, 48, 48, 51, 99}, []byte{})
	//removing unicode escaped chars for >
	x = bytes.ReplaceAll(x, []byte{92, 117, 48, 48, 51, 101, 32}, []byte{})
	//removing // and replacing with newline and a tab
	x = bytes.ReplaceAll(x, []byte{92, 92, 32, 32, 32, 32}, []byte{10, 9})
	//removing tabs
	x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9, 9, 9}, []byte{})
	//removing quotes
	x = bytes.ReplaceAll(x, []byte{34}, []byte{})
	//removing end commas
	x = bytes.ReplaceAll(x, []byte{44, 10}, []byte{10})
	//removing trailing tabs
	x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 9, 9, 9, 93}, []byte{})
	fmt.Printf("%s\n", x)
}

func GetNetIP(j *json.RawMessage) {
	type ipv4 struct {
		System struct {
			Ip struct {
				Ipv4 json.RawMessage `json:"ipv4"`
			} `json:"ip"`
		} `json:"system"`
	}
	var ip ipv4
	err := sonic.Unmarshal(*j, &ip)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	x := ip.System.Ip.Ipv4
	//removing bracket
	x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
	//removing unicode escaped chars for <
	x = bytes.ReplaceAll(x, []byte{92, 117, 48, 48, 51, 99}, []byte{})
	//removing unicode escaped chars for >
	x = bytes.ReplaceAll(x, []byte{92, 117, 48, 48, 51, 101, 32}, []byte{})
	//removing // and replacing with newline and a tab
	x = bytes.ReplaceAll(x, []byte{92, 92, 32, 32, 32, 32}, []byte{10, 9})
	//removing tabs
	x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9, 9, 9}, []byte{})
	//removing quotes
	x = bytes.ReplaceAll(x, []byte{34}, []byte{})
	//removing end commas
	x = bytes.ReplaceAll(x, []byte{44, 10}, []byte{10})
	//removing trailing tabs
	x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 9, 9, 9, 93}, []byte{})
	fmt.Printf("%s\n", x)
	//fmt.Println(ipLinks.System.Ip.Link)
}

func GetNetIptables(j *json.RawMessage) {
	type system struct {
		System struct {
			Iptables struct {
				Filter map[string]json.RawMessage `json:"filter"`
				Mangle map[string]json.RawMessage `json:"mangle"`
				Nat    map[string]json.RawMessage `json:"nat"`
			} `json:"iptables"`
		} `json:"system"`
	}
	var iptables system
	err := sonic.Unmarshal(*j, &iptables)
	if err != nil {
		fmt.Errorf("Cannot unmarshal NodeContents")
	}
	var wg = &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		x := iptables.System.Iptables.Filter[""]
		//removing bracket
		x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
		//removing tabs
		x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9, 9, 9, 9}, []byte{})
		//removing quotes
		x = bytes.ReplaceAll(x, []byte{34}, []byte{})
		//removing end commas
		x = bytes.ReplaceAll(x, []byte{44, 10}, []byte{10})
		//removing trailing tabs
		x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 9, 9, 9, 9, 93}, []byte{})
		fmt.Printf("\n\n----- FILTER TABLE -----\n%s\n", x)

		wg.Done()
	}()
	go func() {
		x := iptables.System.Iptables.Nat[""]
		//removing bracket
		x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
		//removing tabs
		x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9, 9, 9, 9}, []byte{})
		//removing quotes
		x = bytes.ReplaceAll(x, []byte{34}, []byte{})
		//removing end commas
		x = bytes.ReplaceAll(x, []byte{44, 10}, []byte{10})
		//removing trailing tabs
		x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 9, 9, 9, 9, 93}, []byte{})
		fmt.Printf("\n\n----- NAT TABLE -----\n%s\n", x)
		wg.Done()
	}()
	go func() {
		x := iptables.System.Iptables.Mangle[""]
		//removing bracket
		x = bytes.ReplaceAll(x, []byte{91, 10}, []byte{10})
		//removing tabs
		x = bytes.ReplaceAll(x, []byte{9, 9, 9, 9, 9, 9, 9, 9}, []byte{})
		//removing quotes
		x = bytes.ReplaceAll(x, []byte{34}, []byte{})
		//removing end commas
		x = bytes.ReplaceAll(x, []byte{44, 10}, []byte{10})
		//removing trailing tabs
		x = bytes.ReplaceAll(x, []byte{10, 9, 9, 9, 9, 9, 9, 9, 93}, []byte{})
		fmt.Printf("\n\n----- MANGLE TABLE -----\n%s\n", x)
		wg.Done()
	}()
	wg.Wait()
}
