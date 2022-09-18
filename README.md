
# Broker CLI
## About

This repository is the home of the cli used to parse the `dsinfo.json` file from the support bundle generated from Mirantis Kubernetes Engine.

This is tool has been inspired from `docker-cli` and re-used a majority of types from moby,docker-cli repositories.

# Reference
Pre-requisites:
- Set `broker` binary to your bin directory and provide appropriate permission.
- Then change your current directory to the root directory of a support bundle where the `dsinfo.json` file is present.
- Then execute the broker commands.

## broker
To list available commands, either run `broker` with no parameters
or execute `broker help`:
```
A cli tool/parser/analyzer to analyze support bundle's dsinfo.json file and show meaningful output

Usage:
broker [flags]
broker [command]

Available Commands:
completion  Generate the autocompletion script for the specified shell
container   Check containers information
help        Command to show the use of broker commands
info        Shows information about various node specific components
logs        Fetch the logs of the containers and nodes
network     Shows network information of a node
node        Collect Swarm Node information
stats       Shows stats about various node specific components

Flags:
-h, --help   help command

Use "broker [command] --help" for more information about a command.
```
### Command line tree so far
```
broker
├── node [✓]
│   ├── ls
│   └── inspect
├── container [✓]
│   ├── ls
│   └── inspect
├── network [✓]
│   ├── ip
│   ├── links
│   ├── routes
│   └── iptables
├── logs [✓]
│   ├── container
│   └── node
├── info [✓]
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
│   ├── lscpu
│   ├── system-release
│   ├── ucp-controller-diag
│   ├── uptime
│   └── daemondotjson
├── stats [✓]
│   ├── df
│   ├── dockerstat
│   ├── iostatt
│   ├── meminfo
│   ├── netstats
│   └── vmstat
├── Search [ ]
└── Analysis [ ]
```
### node commands
**Usage**
```
╰──➤ broker node  -h
Collect Swarm Node information

Usage:
broker node [COMMAND] [flags]
broker node [command]

Aliases:
node, nodes, nd, node, n

Available Commands:
inspect     Display detailed information on one or more nodes
ls          List swarm nodes

Flags:
-h, --help   help for node

Use "broker node [command] --help" for more information about a command.

```
1. List all nodes
```
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
2. List nodes quietly (with short alias)
```
╰──➤ broker n l -q
ip-172-31-16-116
ip-172-31-26-152
ip-172-31-25-139
.....
```
3. Shows node inspect
```
╰──➤ broker node inspect ip-172-31-16-116
{"ID":"stn47x52941yu9b8752fadmqv","Version":{"Index":972},"CreatedAt":"2022-08-26T10:26:33.976060785Z","UpdatedAt":"2022-08-26T10:44:24.581751062Z","Spec":{"Labels":{"com.docker.ucp.SANs":"ec2-18-159-202-190.eu-central-1.compute.amazonaws.com,ec2-18-185-240-19.eu-central-1.compute.amazonaws.com,ec2-3-67-177-166.eu-central-1.compute.amazonaws.com,ip-172-31-16-116","com.docker.ucp.access.label":"/System","com.docker.ucp.collection":"system","com.docker.ucp.collection.root":"true","com.docker.ucp.collection.swarm":"true","com.docker.ucp.collection.system":"true","com.docker.ucp.node-state-augmented.last-reconciled-time":"2022-08-26T10:44:41Z","com.docker.ucp.node-state-augmented.reconciler-ucp-version":"3.5.5","com.docker.ucp.node-state-augmented.time":"2022-08-26T11:21:58Z","com.docker.ucp.orchestrator.kubernetes":"true","com.docker.ucp.orchestrator.swarm":"true","com.mirantis.launchpad.managed":"true"},"Role":"manager","Availability":"active"},"Description":{"Hostname":"ip-172-31-16-116","Platform":{"Architecture":"x86_64","OS":"linux"},"Resources":{"NanoCPUs":4000000000,"MemoryBytes":7804575744},"Engine":{"EngineVersion":"20.10.7","Labels":{"com.docker.security.apparmor":"enabled","com.docker.security.seccomp":"enabled"},"Plugins":[{"Type":"Log","Name":"awslogs"},{"Type":"Log","Name":"fluentd"},{"Type":"Log","Name":"gcplogs"},{"Type":"Log","Name":"gelf"},{"Type":"Log","Name":"journald"},{"Type":"Log","Name":"json-file"},{"Type":"Log","Name":"local"},{"Type":"Log","Name":"logentries"},{"Type":"Log","Name":"splunk"},{"Type":"Log","Name":"syslog"},{"Type":"Network","Name":"bridge"},{"Type":"Network","Name":"host"},{"Type":"Network","Name":"ipvlan"},{"Type":"Network","Name":"macvlan"},{"Type":"Network","Name":"null"},{"Type":"Network","Name":"overlay"},{"Type":"Volume","Name":"local"}]},"TLSInfo":{"TrustRoot":"-----BEGIN CERTIFICATE-----\nMIIBajCCARCgAwIBAgIUEa7D5Mzk5lvRoCk+LpyIho/paAUwCgYIKoZIzj0EAwIw\nEzERMA8GA1UEAxMIc3dhcm0tY2EwHhcNMjIwODI2MTAyMjAwWhcNNDIwODIxMTAy\nMjAwWjATMREwDwYDVQQDEwhzd2FybS1jYTBZMBMGByqGSM49AgEGCCqGSM49AwEH\nA0IABAb6nZXXDPwj6VFtIOAuPE7+0rUn8OdpCC8yADP0FZemQbs4HN0fNb2bdtDo\n2LjSthD+GfgCdhXnNLYkciLjRfyjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNVHRMB\nAf8EBTADAQH/MB0GA1UdDgQWBBQRQhfvRrsrjJihIMfeTXspqa1QuTAKBggqhkjO\nPQQDAgNIADBFAiEAgm6HRNVVUjlKBf2n7iXLi8FgZWW7eMoJytlFBsMmZPsCICGy\nrJCgzzY83oS+VP/twByv8al83I4XvSLwerQmU7iO\n-----END CERTIFICATE-----\n","CertIssuerSubject":"MBMxETAPBgNVBAMTCHN3YXJtLWNh","CertIssuerPublicKey":"MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBvqdldcM/CPpUW0g4C48Tv7StSfw52kILzIAM/QVl6ZBuzgc3R81vZt20OjYuNK2EP4Z+AJ2Fec0tiRyIuNF/A=="}},"Status":{"State":"ready","Message":"Healthy MKE manager","Addr":"172.31.16.116"},"ManagerStatus":{"Leader":true,"Reachability":"reachable","Addr":"172.31.16.116:2377"}}
```
4. Shows node inspect with pretty flag (with short alias)
```
╰──➤ broker nd ins ip-172-31-16-116 -p
{
"ID": "stn47x52941yu9b8752fadmqv",
"Version": {
"Index": 972
},
"CreatedAt": "2022-08-26T10:26:33.976060785Z",
"UpdatedAt": "2022-08-26T10:44:24.581751062Z",
"Spec": {
"Labels": { .........
.....................
```



### container commands
**Usage**
```
╰──➤ broker container -h
Check containers information

Usage:
broker container [flags]
broker container [command]

Aliases:
container, c, cont, container

Available Commands:
inspect     Display detailed information on one or more containers
ls          List containers

Flags:
-h, --help   help for container

Use "broker container [command] --help" for more information about a command.

```

1. List all containers from all node (not recommended due to heavy mem usage)
```
╰──➤ broker container ps | head -n 10
+-----------------------------------------------------------------------------------------------------------------------------------+
| ip-172-31-26-152                                                                                                                  |
+-------------------------------------+----------------------------+---------------------+---------+-------+--------+---------------+
|                 NAME                |            IMAGE           |       CREATED       |  STATUS |  OOMK | RST-NG |    NETWORK    |
+-------------------------------------+----------------------------+---------------------+---------+-------+--------+---------------+
| k8s_POD_ucp-nvidia-device-partition | ucp-pause:3.5.5            | 2022-08-26T10:29:38 | running | false | false  | host          |
| k8s_ucp-node-feature-discovery-work | sha256:56244037            | 2022-08-26T11:21:30 | exited  | false | false  |               |
| ucp-hardware-info                   | ucp-hardware-info:3.5.5    | 2022-08-26T10:29:32 | running | false | false  | ucp-bridge    |
```
2. List containers from specific node (using short alias) [for full name of the container use the flag `-v`]
```
╰──➤ broker c l ip-172-31-26-152
+-----------------------------------------------------------------------------------------------------------------------------------+
| ip-172-31-26-152                                                                                                                  |
+-------------------------------------+----------------------------+---------------------+---------+-------+--------+---------------+
|                 NAME                |            IMAGE           |       CREATED       |  STATUS |  OOMK | RST-NG |    NETWORK    |
+-------------------------------------+----------------------------+---------------------+---------+-------+--------+---------------+
| k8s_POD_calico-node-bt469_kube-syst | ucp-pause:3.5.5            | 2022-08-26T10:29:38 | running | false | false  | host          |
| ucp-kube-proxy                      | ucp-hyperkube:3.5.5        | 2022-08-26T10:29:34 | running | false | false  | host          |
| ucp-swarm-manager                   | ucp-swarm:3.5.5            | 2022-08-26T10:29:39 | running | false | false  | ucp-bridge    |
| k8s_ucp-node-feature-discovery-mast | sha256:56244037            | 2022-08-26T11:21:30 | exited  | false | false  |               |
.......................
```
3. Show container inspect [using `-p` for pretty output],
```
╰──➤ broker c inspect ucp-kv ip-172-31-26-152 -p
{
"Id": "772190ac5856455ae2df885028de34fbbadb4ac1aa7f7f797fdebef504d68f33",
"Created": "2022-08-26T10:29:15.352759199Z",
"Path": "/bin/entrypoint.sh",
"Args": [
"--data-dir",
"/data/datav3",
"--name",
"orca-kv-172.31.26.152",
"--listen-peer-urls",
"https://0.0.0.0:12380",
"--listen-client-urls",
```
### network commands
**Usage**
```
╰──➤ broker network -h
Shows network information of a node

Usage:
broker network [argument] [nodeName]

Allowed arguments are:
ip
link
route
iptables [flags]

Aliases:
network, net, network, nt

Flags:
-h, --help   help for network
```
1. Show `ifconfig` like output from a node
```
╰──➤ broker network ip ip-172-31-16-116
1: lo    inet 127.0.0.1/8 scope host lo
valid_lft forever preferred_lft forever
2: ens3    inet 172.31.16.116/20 brd 172.31.31.255 scope global dynamic ens3
valid_lft 3264sec preferred_lft 3264sec
3: docker0    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
valid_lft forever preferred_lft forever
12: docker_gwbridge    inet 172.18.0.1/16 brd 172.18.255.255 scope global docker_gwbridge
valid_lft forever preferred_lft forever
61: br-b1811fd9579a    inet 172.19.0.1/16 brd 172.19.255.255 scope global br-b1811fd9579a
valid_lft forever preferred_lft forever
92: vxlan.calico    inet 192.168.127.128/32 scope global vxlan.calico
valid_lft forever preferred_lft forever
```
2. Shows all links of a node
```
╰──➤ broker net link ip-172-31-16-116
1: lo: LOOPBACK,UP,LOWER_UPmtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
2: ens3: BROADCAST,MULTICAST,UP,LOWER_UPmtu 9001 qdisc mq state UP mode DEFAULT group default qlen 1000
link/ether 02:29:38:c6:c9:ee brd ff:ff:ff:ff:ff:ff
altname enp0s3
3: docker0: BROADCAST,MULTICAST,UP,LOWER_UPmtu 1500 qdisc noqueue state UP mode DEFAULT group default
link/ether 02:42:a3:e3:32:d3 brd ff:ff:ff:ff:ff:ff
12: docker_gwbridge: BROADCAST,MULTICAST,UP,LOWER_UPmtu 1500 qdisc noqueue state UP mode DEFAULT group default
link/ether 02:42:86:c4:89:9d brd ff:ff:ff:ff:ff:ff
14: veth374801c@if13: BROADCAST,MULTICAST,UP,LOWER_UPmtu 1500 qdisc noqueue master docker_gwbridge state UP mode DEFAULT group default
link/ether 7a:d1:d6:bd:8e:0a brd ff:ff:ff:ff:ff:ff link-netnsid 1
61: br-b1811fd9579a: BROADCAST,MULTICAST,UP,LOWER_UPmtu 1500 qdisc noqueue state UP mode DEFAULT group default
link/ether 02:42:6b:c0:05:a4 brd ff:ff:ff:ff:ff:ff
63: veth12e8ea2@if62: BROADCAST,MULTICAST,UP,LOWER_UPmtu 1500 qdisc noqueue master br-b1811fd9579a state UP mode DEFAULT group default
........................
```
3. Shows route of a node
```
╰──➤ broker net route ip-172-31-16-116
default via 172.31.16.1 dev ens3 proto dhcp src 172.31.16.116 metric 100
172.17.0.0/16 dev docker0 proto kernel scope link src 172.17.0.1
172.18.0.0/16 dev docker_gwbridge proto kernel scope link src 172.18.0.1
172.19.0.0/16 dev br-b1811fd9579a proto kernel scope link src 172.19.0.1
172.31.16.0/20 dev ens3 proto kernel scope link src 172.31.16.116
172.31.16.1 dev ens3 proto dhcp scope link src 172.31.16.116 metric 100
192.168.72.0/26 via 192.168.72.0 dev vxlan.calico onlink
192.168.90.192/26 via 192.168.90.192 dev vxlan.calico onlink
192.168.105.64/26 via 192.168.105.64 dev vxlan.calico onlink
192.168.127.129 dev cali57771604549 scope link
192.168.162.192/26 via 192.168.162.193 dev vxlan.calico onlink
192.168.164.128/26 via 192.168.164.129 dev vxlan.calico onlink
192.168.240.64/26 via 192.168.240.65 dev vxlan.calico onlink
192.168.245.128/26 via 192.168.245.128 dev vxlan.calico onlink
```
4. Show iptables of a node
```
╰──➤ broker net iptables ip-172-31-16-116
----- MANGLE TABLE -----
Chain cali-to-host-endpoint (1 references)
pkts bytes target     prot opt in     out     source               destination

----- FILTER TABLE -----
Chain cali-wl-to-host (1 references)
pkts bytes target     prot opt in     out     source               destination
9166  946K cali-from-wl-dispatch  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cali:Ee9Sbo10IpVujdIY */
77  4620 ACCEPT     all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cali:nSZbcOoG1xPONxb8 */ /* Configured DefaultEndpointToHostAction */

----- NAT TABLE -----
Chain cali-nat-outgoing (1 references)
pkts bytes target     prot opt in     out     source               destination
663 61346 MASQUERADE  all  --  *      *       0.0.0.0/0            0.0.0.0/0            /* cali:flqWnvo8yq4ULQLa */ match-set cali40masq-ipam-pools src ! match-set cali40all-ipam-pools dst random-fully
```
### logs commands
**Usage**
```
╰──➤ broker logs -h
Fetch the logs of the containers and nodes

Usage:
broker logs [node|container] [nodeName|containerName] [flags]
broker logs [command]

Available Commands:
container   Fetch the logs of the containers from specific node
node        Fetch the node specific logs

Flags:
-h, --help   help for logs

Use "broker logs [command] --help" for more information about a command.
```
1. List container logs
```
╰──➤ broker logs container ucp-kv ip-172-31-26-152
2022-08-26T10:29:15.987674514Z cp: recursion detected, omitting directory '/data/datav3'
2022-08-26T10:29:15.988533706Z Joining existing etcd cluster ...
2022-08-26T10:29:15.988547370Z Waiting for local peer https://172.31.26.152:12380 to appear in etcd cluster membership
2022-08-26T10:29:16.027071317Z ...
2022-08-26T10:29:21.058611155Z ...
2022-08-26T10:29:26.076289996Z ...
2022-08-26T10:29:31.111532697Z Starting etcd server with arguments: --data-dir /data/datav3 --name orca-kv-172.31.26.152 --listen-peer-urls https://0.0.0.0:12380 --listen-client-urls https://0.0.0.0:2379 --advertise-client-urls https://172.31.26.152:12379 --trusted-ca-file /etc/docker/ssl/ca.pem --cert-file /etc/docker/ssl/cert.pem --key-file /etc/docker/ssl/key.pem --client-cert-auth --peer-trusted-ca-file /etc/docker/ssl/ca.pem --peer-cert-file /etc/docker/ssl/cert.pem --peer-key-file /etc/docker/ssl/key.pem --peer-client-cert-auth --heartbeat-interval 500 --election-timeout 5000 --snapshot-count 20000 --enable-v2 --initial-cluster-state existing --initial-cluster orca-kv-172.31.26.152=https://172.31.26.152:12380,orca-kv-172.31.16.116=https://172.31.16.116:12380 --initial-advertise-peer-urls https://172.31.26.152:12380 --metrics=extensive --logger=zap
```
2. The logs available for a node,
```
╰──➤ broker logs node -h
Fetch the node specific logs
Usage:
broker logs node [arguments] [nodeName]

Allowed arguments are:
calico
deadcontainermounts
dmesg
dtr
daemon
kernel
shimlogs [flags]

Flags:
-h, --help   help for node
```
3. Checking docker daemon logs for a node
```
╰──➤ broker logs node daemon ip-172-31-26-152 | head
time="2022-08-26T11:20:04.982599591Z" level=debug msg="Calling GET /v1.40/containers/json?all=1\u0026filters=%7B%22id%22%3A%7B%22148b633f248f0de23e1a049825ed58d2e0f10f00c188a17ddfc19c07925ec3e2%22%3Atrue%7D%7D\u0026limit=0"
time="2022-08-26T11:20:04.982651328Z" level=debug msg="attach: stderr: begin"
time="2022-08-26T11:20:04.982743811Z" level=debug msg="attach: stdout: begin"
time="2022-08-26T11:20:04.982625077Z" level=debug msg="Calling GET /v1.40/containers/json?all=1\u0026filters=%7B%22id%22%3A%7B%22148b633f248f0de23e1a049825ed58d2e0f10f00c188a17ddfc19c07925ec3e2%22%3Atrue%7D%7D\u0026limit=0"
time="2022-08-26T11:20:04.985751049Z" level=debug msg=event module=libcontainerd namespace=moby topic=/tasks/exec-added
time="2022-08-26T11:20:05.039440575Z" level=debug msg="attach: stdout: end"
```
4. Checking kernel logs for a node
```
╰──➤ broker logs node kernel ip-172-31-26-152 | head
2038-01-19T09:14:07+06:00   5   kernel    ubuntu    Linux version 5.15.0-1017-aws (buildd@lcy02-amd64-055) (gcc (Ubuntu 9.4.0-1ubuntu1~20.04.1) 9.4.0, GNU ld (GNU Binutils for Ubuntu) 2.34) #21~20.04.1-Ubuntu SMP Fri Aug 5 11:44:14 UTC 2022 (Ubuntu 5.15.0-1017.21~20.04.1-aws 5.15.39)
2038-01-19T09:14:07+06:00   6   kernel    ubuntu    Command line: BOOT_IMAGE=/boot/vmlinuz-5.15.0-1017-aws root=PARTUUID=2a294749-f511-4cc3-84af-dd528794db24 ro console=tty1 console=ttyS0 nvme_core.io_timeout=4294967295 panic=-1
2038-01-19T09:14:07+06:00   6   kernel    ubuntu    KERNEL supported cpus:
2038-01-19T09:14:07+06:00   6   kernel    ubuntu      Intel GenuineIntel
2038-01-19T09:14:07+06:00   6   kernel    ubuntu      AMD AuthenticAMD
2038-01-19T09:14:07+06:00   6   kernel    ubuntu      Hygon HygonGenuine
2038-01-19T09:14:07+06:00   6   kernel    ubuntu      Centaur CentaurHauls
2038-01-19T09:14:07+06:00   6   kernel    ubuntu      zhaoxin   Shanghai
2038-01-19T09:14:07+06:00   6   kernel    ubuntu    x86/fpu: Supporting XSAVE feature 0x001: 'x87 floating point registers'
2038-01-19T09:14:07+06:00   6   kernel    ubuntu    x86/fpu: Supporting XSAVE feature 0x002: 'SSE registers'
```
5. Checking shimlogs from a node
```
╰──➤ broker logs node shimlogs ip-172-31-26-152 | head
-----------------------------------------------------
ucp-kubelet-86250a63b5cc-shim-debug.log
-----------------------------------------------------
rootfs: [\u0026Mount{Type:overlay,Source:overlay,Target:,Options:[index=off workdir=/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/21/work upperdir=/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/21/fs lowerdir=/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/19/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/18/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/17/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/16/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/15/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/14/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/13/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/12/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/11/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/10/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/9/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/8/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/7/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/6/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/5/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/4/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/3/fs:/var/lib/containerd/io.containerd.snapshotter.v1.overlayfs/snapshots/2/fs],XXX_unrecognized:[],}]
resolveRootfsCommand: /run/containerd/io.containerd.runtime.v2.task/com.docker.ucp/ucp-kubelet-86250a63b5cc/rootfs/bin/kubelet_entrypoint.sh
waiting for command exit event
setting up IO
```
### info commands
**Usage**
```
╰──➤ broker info -h
Shows information about various node specific components

Usage:
broker info [argument] [nodeName]

Allowed arguments are:
lscpu
clusterinfo
dmidecode
images
dockerinfo
pidlimits
version
filemax
ifconfig
kernel
mounts
networking
psauxdocker
rethinkstatus
running-cgroup
sestatus
kernelconf
ssd
system-cgroup
system-release
ucp-controller-diag
uptime
daemondotjson [flags]

Aliases:
info, info, i, inf

Flags:
-h, --help      help for info
-p, --pretty    pretty JSON output
-v, --verbose   verbose output
```
1. List info from all of the node. NOT RECOMMENDED for heavy memory usage. Instead use something like `for i in $(broker n l -q); do broker info kernel $i; done)`
```
╰──➤ broker info kernel
Node : ip-172-31-28-185
Linux version 5.15.0-1017-aws (buildd@lcy02-amd64-055) (gcc (Ubuntu 9.4.0-1ubuntu1~20.04.1) 9.4.0, GNU ld (GNU Binutils for Ubuntu) 2.34) #21~20.04.1-Ubuntu SMP Fri Aug 5 11:44:14 UTC 2022
Node : ip-172-31-16-116
Linux version 5.15.0-1017-aws (buildd@lcy02-amd64-055) (gcc (Ubuntu 9.4.0-1ubuntu1~20.04.1) 9.4.0, GNU ld (GNU Binutils for Ubuntu) 2.34) #21~20.04.1-Ubuntu SMP Fri Aug 5 11:44:14 UTC 2022
............
```
2. Show version info of a specific node
```
╰──➤ broker info version ip-172-31-28-83
Client:
Version: 20.10.12
API Version: 1.41
Server:
Version: 20.10.7
API Version: 1.41
License : Mirantis Container Runtime (this node is not a swarm manager - check license status on a manager node)
Kernel:
Version: 5.15.0-1017-aws
```
3. Show daemon.json file contents of a node
```
╰──➤ broker i daemondotjson ip-172-31-16-116 -p
{
"debug": true,
"log-opts": {
"max-file": "3",
"max-size": "10m"
}
}
```
4. Show lsb_release info from a node
```
╰──➤ broker info system-release ip-172-31-16-116
NAME=Ubuntu
VERSION=20.04.4 LTS (Focal Fossa)
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME=Ubuntu 20.04.4 LTS
VERSION_ID=20.04
HOME_URL=https://www.ubuntu.com/
SUPPORT_URL=https://help.ubuntu.com/
BUG_REPORT_URL=https://bugs.launchpad.net/ubuntu/
PRIVACY_POLICY_URL=https://www.ubuntu.com/legal/terms-and-policies/privacy-policy
VERSION_CODENAME=focal
UBUNTU_CODENAME=focal"
```
5. Shows `lscpu` output of a node
```
╰──➤ broker info lscpu ip-172-31-16-116
Architecture:                    x86_64
CPU op-mode(s):                  32-bit, 64-bit
Address sizes:                   46 bits physical, 48 bits virtual
Byte Order:                      Little Endian
CPU(s):                          4
On-line CPU(s) list:             0-3
Vendor ID:                       GenuineIntel
Model name:                      Intel(R) Xeon(R) CPU E5-2666 v3 @ 2.90GHz
..................
```
Feel free to explore other arguments too. Some of the arguments might return nil because it wasn't able to collect those information.


### stats commands
**Usage**
```
╰──➤ broker stats -h
Shows stats about various node specific components

Usage:
broker stats [argument] [nodeName]

Allowed arguments are:
df
dockerstats
iostat
meminfo
netstat
vmstat [flags]

Flags:
-h, --help   help for stats
```
1. Shows disk free (`df`) of a node
```
╰──➤ broker stats df ip-172-31-16-116
Filesystem      Size  Used Avail Use% Mounted on
overlay          49G   10G   39G  21% /
tmpfs            64M     0   64M   0% /dev
shm              64M     0   64M   0% /dev/shm
/dev/root        49G   10G   39G  21% /etc
/dev/xvda15     105M  5.2M  100M   5% /boot/efi
tmpfs           745M  4.3M  741M   1% /run
tmpfs           5.0M     0  5.0M   0% /run/lock
overlay          49G   10G   39G  21% /run/containerd/io.containerd.runtime.v2.task/com.docker.ucp/ucp-kubelet-fe2e226f482b/rootfs
.....................
```
2. Shows docker stats from a node,
```
╰──➤ broker stats dockerstats ip-172-31-16-116
+-------------------------------------+-----+---------+---------+---------------------+-----------------+-----------------+
|                 NAME                | PID | CPUPERC | MEMPERC |       MEMUSAGE      |     BLOCKIO     | NETIO           |
+-------------------------------------+-----+---------+---------+---------------------+-----------------+-----------------+
| romantic_einstein                   | 8   | 37.57%  | 0.22%   | 16.7MiB / 7.269GiB  | 0B / 365kB      | 0B / 0B         |
| ucp-manager-agent.stn47x52941yu9b87 | 8   | 0.00%   | 0.28%   | 21.17MiB / 7.269GiB | 135kB / 2.78MB  | 1.49MB / 1.68MB |
| ......................................................................................................................  |
| ......................................................................................................................  |
| ucp-kube-proxy                      | 7   | 0.00%   | 0.16%   | 12.23MiB / 7.269GiB | 860kB / 0B      | 0B / 0B         |
| ucp-kubelet                         | 19  | 6.95%   | 0.30%   | 22.69MiB / 7.269GiB | 1.8MB / 0B      | 1.19kB / 0B     |
| ucp-kube-scheduler                  | 9   | 1.64%   | 0.30%   | 22.51MiB / 7.269GiB | 1.3MB / 0B      | 9.27MB / 1.79MB |
| ucp-kube-controller-manager         | 8   | 1.76%   | 0.77%   | 57.28MiB / 7.269GiB | 954kB / 0B      | 23.6MB / 6.1MB  |
| ucp-kube-apiserver                  | 18  | 14.60%  | 4.25%   | 316MiB / 7.269GiB   | 778kB / 12.3kB  | 350MB / 223MB   |
| ucp-controller                      | 17  | 13.66%  | 2.28%   | 170.1MiB / 7.269GiB | 3.23MB / 2.96MB | 563MB / 332MB   |
| ucp-swarm-manager                   | 8   | 6.62%   | 0.50%   | 37.37MiB / 7.269GiB | 0B / 0B         | 521MB / 518MB   |
| ucp-proxy                           | 9   | 2.90%   | 0.37%   | 27.63MiB / 7.269GiB | 65.5kB / 0B     | 87.6MB / 427MB  |
| ucp-auth-api.stn47x52941yu9b8752fad | 11  | 0.00%   | 0.19%   | 14.44MiB / 7.269GiB | 811kB / 0B      | 0B / 0B         |
| ucp-auth-store                      | 82  | 0.14%   | 1.90%   | 141.8MiB / 7.269GiB | 2.16MB / 31.3MB | 124MB / 317MB   |
| ucp-kv                              | 13  | 5.59%   | 1.15%   | 85.24MiB / 7.269GiB | 2.31MB / 409MB  | 118MB / 430MB   |
+-------------------------------------+-----+---------+---------+---------------------+-----------------+-----------------+
```
3. Shows `meminfo` of all nodes. [you can also use `broker stats meminfo` without any node name but following approach is recommended]
```
╰──➤ for i in $(broker n l -q); do printf "\nNode Name : $i\n" ;broker stats meminfo $i | head -n 5; done
Node Name : ip-172-31-28-83
MemTotal:        7621660 kB
MemFree:          350232 kB
MemAvailable:    6432016 kB
Buffers:          147148 kB

Node Name : ip-172-31-25-139
MemTotal:        7621660 kB
MemFree:          162152 kB
MemAvailable:    6029596 kB
Buffers:          204852 kB

Node Name : ip-172-31-22-22
MemTotal:        7621656 kB
MemFree:          175412 kB
MemAvailable:    6063976 kB
Buffers:          208168 kB
```
4. Shows `iostat`
```
╰──➤ broker stats iostat ip-172-31-26-152
Linux 5.15.0-1017-aws (ip-172-31-26-152) \t08/26/22 \t_x86_64_\t(4 CPU)
08/26/22 11:22:25
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
7.98    0.05    4.61    0.77    0.32   86.28

Device            r/s     rkB/s   rrqm/s  %rrqm r_await rareq-sz     w/s     wkB/s   wrqm/s  %wrqm w_await wareq-sz     d/s     dkB/s   drqm/s  %drqm d_await dareq-sz     f/s f_await  aqu-sz  %util
loop0            0.09      2.46     0.00   0.00    0.79    28.65    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.03
loop1            0.01      0.13     0.00   0.00    2.56     8.86    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.00
loop2            0.08      0.65     0.00   0.00    1.12     8.57    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.01
loop3            0.02      0.32     0.00   0.00    2.16    15.46    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.00
loop4            0.19      6.78     0.00   0.00    0.60    35.52    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.07
loop5            0.00      0.00     0.00   0.00    0.09     1.27    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.00
xvda             3.55    130.30     1.24  25.90    2.02    36.73   96.57   3648.83   123.33  56.08    5.53    37.78    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.54   9.94

08/26/22 11:22:26
avg-cpu:  %user   %nice %system %iowait  %steal   %idle
6.39    0.00    4.35    0.26    0.00   89.00

Device            r/s     rkB/s   rrqm/s  %rrqm r_await rareq-sz     w/s     wkB/s   wrqm/s  %wrqm w_await wareq-sz     d/s     dkB/s   drqm/s  %drqm d_await dareq-sz     f/s f_await  aqu-sz  %util
loop0            0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.00
loop1            0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.00
loop2            0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.00
loop3            0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.00
loop4            0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.00
loop5            0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.00   0.00
xvda             0.00      0.00     0.00   0.00    0.00     0.00   26.00    420.00    74.00  74.00    0.81    16.15    0.00      0.00     0.00   0.00    0.00     0.00    0.00    0.00    0.02   5.60
```