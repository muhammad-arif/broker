# Broker CLI
## About

This repository is the home of the cli used to parse the `dsinfo.json` file from the support bundle generated from Mirantis Kubernetes Engine.

This is tool has been inspired from `docker-cli` and re-used a majority of types from moby,docker-cli repositories.
# Reference
## broker
To list available commands, either run `broker` with no parameters
or execute `broker help`:
```console=
A cli tool parser/analyzer to analyze support bundle's dsinfo.json file and show meaningful output

Usage:  broker [OPTIONS] COMMAND

Available Commands:
node	      Shows node specific information
container   Shows container specific information
logs        Shows logs of a container or a specific object of a node
info        Shows components related information of the nodes and container
stats	      Shows statistics/metrics/performance related information

Flags:
-h, --help  Help for Broker
```
### Command line tree so far
```
broker
├── node
│   ├── ls
│   └── inspect
├── container
│   ├── ls
│   └── inspect
├── logs
│   ├── container
│   └── node
├── info
│   ├── clusterinfo
│   ├── dmidecode
│   ├── images
│   ├── dockerinfo
│   ├── pidlimits
│   ├── version
│   ├── filemax
│   ├── ifconfig
│   ├── kernelversion
│   ├── mounts
│   ├── networking
│   ├── psauxdocker
│   ├── rethinkstatus
│   ├── running-cgroup
│   ├── sestatus
│   ├── kernelconfig
│   ├── ssd
│   ├── system-cgroup
│   ├── system-release
│   ├── ucp-controller-diag
│   ├── uptime
│   └── daemondotjson
├── stats
│   ├── df
│   ├── dockerstat
│   ├── iostatt
│   ├── lscpu
│   ├── meminfo
│   ├── netstats
│   └── vmstat
├── Search
└── Analysis
```
### node
The command `node` provides 2 subcommands `ls` and `inspect`
#### node ls
This command will return list of all the nodes of the cluster,
```console
╰──➤ broker node ls
+------------------+---------+--------+-------+---------------+---------+---------+-------+-------+
|       NAME       |   ROLE  |  AVAIL | STATE |       IP      |    OS   |  ENGINE |  MKE  |  ORCA |
+------------------+---------+--------+-------+---------------+---------+---------+-------+-------+
| ip-172-31-16-116 | LEADER  | active | ready | 172.31.16.116 | linux   | 20.10.7 | 3.5.5 | mixed |
| ip-172-31-26-152 | manager | active | ready | 172.31.26.152 | linux   | 20.10.7 | 3.5.5 | mixed |
| ip-172-31-16-152 | manager | active | ready | 172.31.16.152 | linux   | 20.10.7 | 3.5.5 | mixed |
| ip-172-31-28-83  | worker  | active | ready | 172.31.28.83  | linux   | 20.10.7 | 3.5.5 | swarm |
| ip-172-31-25-139 | worker  | active | ready | 172.31.25.139 | linux   | 20.10.7 | 3.5.5 | mixed |
| EC2AMAZ-88CUOUJ  | worker  | active | ready | 172.31.18.27  | windows | 20.10.7 | 3.5.5 | swarm |
| ip-172-31-22-22  | worker  | active | ready | 172.31.22.22  | linux   | 20.10.7 | 3.5.5 | mixed |
| ip-172-31-28-195 | worker  | active | ready | 172.31.28.195 | linux   | 20.10.7 | 3.5.5 | mixed |
| ip-172-31-28-185 | worker  | active | ready | 172.31.28.185 | linux   | 20.10.7 | 3.5.5 | swarm |
| ip-172-31-16-217 | worker  | active | ready | 172.31.16.217 | linux   | 20.10.7 | 3.5.5 | swarm |
+------------------+---------+--------+-------+---------------+---------+---------+-------+-------+
```
For help use help flag,
```console
╰──➤ broker node ls --help
List swarm nodes

Usage:
broker node broker node ls [flags]

Aliases:
ls, l, list

Flags:
-h, --help   help for broker

```
#### node inspect
The `broker node inspect` command shows detailed information on one or more nodes.
Usage,
```
Display detailed information on one or more nodes

Usage:
broker node [OPTIONS] inspect [NODE-NAME] [flags]

Aliases:
ins, i, inspect

Flags:
-h, --help     help for node
-p, --pretty   pretty JSON output

```
Example,
```
╰──➤ broker node inspect ip-172-31-16-152 -p | head
{
"ID": "zxnkpm1ip1jw3y0d2hjftx31p",
"Version": {
"Index": 974
},
"CreatedAt": "2022-08-26T10:28:40.628109733Z",
"UpdatedAt": "2022-08-26T10:44:24.608944111Z",
"Spec": {
"Labels": {
"com.docker.ucp.SANs": "ec2-18-159-202-190.eu-central-1.compute.amazonaws.com,ec2-18-185-240-19.eu-central-1.compute.amazonaws.com,ec2-3-67-177-166.eu-central-1.compute.amazonaws.com",

```
### container
This command returns the information about containers.

This command has 2 subcommand
#### list
The command `broker ps` or `broker container list` will return the all the containers in a specific node if the node name is mentioned. If not it will return all the containers in the cluster node-by-node.
Usage,
```
╰──➤ broker ps --help
List containers

Usage:
broker container ls [all|NODE NAME] [flags]

Aliases:
list, ls, l, ps

Flags:
-h, --help      help for container
-v, --verbose   verbose output


```
Example,
```
╰──➤ broker ps ip-172-31-28-195
+-------------------------------------------------------------------------------------------------------------------------------------------+
| ip-172-31-28-195                                                                                                                          |
+-------------------------------------+---------------------------------------+---------------------+---------+-------+--------+------------+
|                 NAME                |                 IMAGE                 |       CREATED       |  STATUS |  OOMK | RST-NG |   NETWORK  |
+-------------------------------------+---------------------------------------+---------------------+---------+-------+--------+------------+
| dtr-notary-signer-000000000001      | dtr-notary-signer:2.9.5               | 2022-08-26T10:37:20 | running | false | false  | dtr-ol     |
| dtr-scanningstore-000000000001      | dtr-postgres:2.9.5                    | 2022-08-26T10:37:27 | running | false | false  | dtr-ol     |
| k8s_calico-node_calico-node-pcdq7_k | ucp-calico-node@sha256:6b0a9          | 2022-08-26T10:31:10 | running | false | false  |            |
| k8s_install-cni_calico-node-pcdq7_k | ucp-calico-cni@sha256:845ab           | 2022-08-26T10:31:01 | exited  | false | false  |            |
| k8s_ucp-node-feature-discovery-work | sha256:56244037                       | 2022-08-26T11:17:46 | exited  | false | false  |            |
| ucp-kube-proxy                      | ucp-hyperkube:3.5.5                   | 2022-08-26T10:30:32 | running | false | false  | host       |
| ucp-kubelet                         | ucp-agent:3.5.5                       | 2022-08-26T10:30:27 | running | false | false  | ucp-bridge |
| dtr-garant-000000000001             | dtr-garant:2.9.5                      | 2022-08-26T10:36:44 | running | false | false  | dtr-ol     |
| dtr-jobrunner-000000000001          | dtr-jobrunner:2.9.5                   | 2022-08-26T10:37:12 | running | false | false  | dtr-ol     |
| dtr-notary-server-000000000001      | dtr-notary-server:2.9.5               | 2022-08-26T10:36:59 | running | false | false  | dtr-ol     |
| k8s_POD_ucp-node-feature-discovery- | ucp-pause:3.5.5                       | 2022-08-26T10:31:10 | running | false | false  | none       |
| k8s_POD_ucp-nvidia-device-partition | ucp-pause:3.5.5                       | 2022-08-26T10:30:45 | running | false | false  | host       |
| k8s_firewalld-policy_calico-node-pc | sha256:f8f7c56d                       | 2022-08-26T10:30:52 | exited  | false | false  |            |
| k8s_ucp-node-feature-discovery-mast | sha256:56244037                       | 2022-08-26T11:17:45 | exited  | false | false  |            |
| ucp-interlock-config.ajk2iiadewfugd | ucp-interlock-config:3.5.5            | 2022-08-26T10:44:35 | running | false | false  | bridge     |
| ucp-proxy                           | ucp-agent:3.5.5                       | 2022-08-26T10:29:31 | running | false | false  | ucp-bridge |
| dtr-registry-000000000001           | dtr-registry:2.9.5                    | 2022-08-26T10:36:37 | running | false | false  | dtr-ol     |
| k8s_POD_calico-node-pcdq7_kube-syst | ucp-pause:3.5.5                       | 2022-08-26T10:30:45 | running | false | false  | host       |
| k8s_ucp-nvidia-device-partitioner_u | ucp-nvidia-device-plugin@sha256:78375 | 2022-08-26T10:30:57 | exited  | false | false  |            |
| ucp-hardware-info                   | ucp-hardware-info:3.5.5               | 2022-08-26T10:29:34 | running | false | false  | ucp-bridge |
| dtr-api-000000000001                | dtr-api:2.9.5                         | 2022-08-26T10:36:51 | running | false | false  | dtr-ol     |
| dtr-nginx-000000000001              | dtr-nginx:2.9.5                       | 2022-08-26T10:37:08 | running | false | false  | dtr-ol     |
| dtr-rethinkdb-000000000001          | dtr-rethink:2.9.5                     | 2022-08-26T10:36:30 | running | false | false  | dtr-ol     |
| k8s_ucp-pause_ucp-nvidia-device-par | sha256:0bdf787a                       | 2022-08-26T10:30:57 | running | false | false  |            |
| ucp-interlock-config.ajk2iiadewfugd | ucp-interlock-config:3.5.5            | 2022-08-26T10:44:32 | exited  | false | false  | bridge     |
| ucp-worker-agent-x.ajk2iiadewfugdj0 | ucp-agent:3.5.5                       | 2022-08-26T10:29:30 | running | false | false  | bridge     |
+-------------------------------------+---------------------------------------+---------------------+---------+-------+--------+------------+

```
#### inspect
The command `broker container inspect` or `broker inspect` returns similar output of `docker inspect` which is the the detailed information of a particular container.

Usage,
```
╰──➤ broker inspect --help
Display detailed information on one or more containers

Usage:
broker container inspect [NODE NAME] [CONTAINER NAME] [flags]

Aliases:
container, inspect, ins, describe, i

Flags:
-h, --help     help for container
-p, --pretty   pretty JSON output

```

Example,
```
╰──➤ broker inspect ip-172-31-28-195 dtr-rethinkdb-000000000001 -p | head
{
"Id": "c00363644350a0ac5df8647c2fc9d7c56d5ac293298c48ef135e5b7985dd8a50",
"Created": "2022-08-26T10:36:30.069032067Z",
"Path": "/bin/rethinkwrapper",
"Args": [],
"State": {
"Status": "running",
"Running": true,
"Paused": false,
"Restarting": false,

```
### logs
This command returns the logs of the nodes and containers of the cluster. It has 2 subcommands.

#### container
The command `broker logs container` returns the log of specific container of a specific node.

#### nodes