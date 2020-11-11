# lsmods
![goreport passing](https://goreportcard.com/badge/github.com/bieli/lsmods)
[![Build Status](https://travis-ci.org/bieli/lsmods.png)](https://travis-ci.org/bieli/lsmods)

## How to compile
```bash
$ go build -i main.go
```

## How to run
```bash
$ chmod a+x ./main
$ ./main | head
2020/09/28 01:09:31 Currently loaded kernel modules with descriptions:
veth				Virtual Ethernet Tunnel
usblp				USB Printer Device Class driver
rfcomm				Bluetooth RFCOMM ver 1.11
xt_conntrack			Xtables: connection tracking state match
xt_MASQUERADE			Xtables: automatic-address SNAT
nf_conntrack_netlink		
xfrm_user			
xfrm_algo			
xt_addrtype			Xtables: address type match

```




## How to run with Docker


### Build Docker container
```bash
$ sudo docker build -t lsmods .
```

### Debug inside Docker container
```bash
$ docker run -it lsmods /bin/bash
```
