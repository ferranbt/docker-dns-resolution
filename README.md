# Docker custom DNS resolution

This repo describes how to run a custom DNS resolution system for `Docker` containers that does not require any sudo application (outside `Docker` itself).

By default, `Docker` provides a DNS resolution mechanism in which `Docker` containers with the `--name <name>` flag can be reached by its name from other `Docker` containers. However, this utility only works for simple use cases.

Deploy `coredns` service:

```
$ docker run --name coredns --restart=always --volume=$PWD/coredns/:/root/ coredns/coredns -conf /root/Corefile
```

Get the ip of the `Dns` server:

```
$ docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' coredns
172.17.0.2
```

Deploy the `Grpc` resolution service:

```
$ go run main.go
```

Deploy a `Docker` container connected to the DNS service and make a DNS query:

```
$ docker run --dns 172.17.0.2 -it tutum/dnsutils /bin/sh
# dig example.com
```

The console of the `Grpc` resolution backend should print the `DNS` query to be resolved.
