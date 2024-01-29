# lsmods

Terminal tool to show all enabled Linux Kernel modules with descriptions (line by line).

![CI status](https://github.com/bieli/lsmods/actions/workflows/ci.yaml/badge.svg)

![go report](https://goreportcard.com/badge/github.com/bieli/lsmods)


## How to compile
```bash
$ make build
```

## How to run
```bash
$ chmod a+x ./lsmods
$ ./lsmods | tail
2020/11/11 23:17:30 Currently loaded kernel modules with descriptions:
xor				
xt_CHECKSUM			Xtables: checksum modification
xt_MASQUERADE			Xtables: automatic-address SNAT
xt_addrtype			Xtables: address type match
xt_comment			Xtables: No-op match which can be tagged with a comment
xt_conntrack			Xtables: connection tracking state match
xt_nat				
xt_state			ip[6]_tables connection tracking state match module
xt_tcpudp			Xtables: TCP, UDP and UDP-Lite match
zstd_compress			Zstd Compressor
```

## How to run unit test
```bash
$ make test
```

## How to run with Docker

### Build Docker container
```bash
$ sudo docker build -t lsmods .
```

### Debug inside Docker container
Warning ! Below command show kernel modules from **host machine** - not from docker container (it's impossible to do this way).
```bash
$ docker run --privileged --cap-add=ALL -v /dev:/dev -v /lib/modules:/lib/modules -it lsmods /go/src/app/lsmods
```
