module github.com/hfyeh/agent

go 1.21.8

require (
	github.com/docker/docker v1.13.1
	github.com/gogo/protobuf v1.2.0
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/yamux v0.0.0-20180826203732-cc6d2ea263b2
	github.com/mdlayher/vsock v0.0.0-20190429153235-7b7533a7ca4e
	github.com/opencontainers/runc v1.0.0-rc9.0.20200122160610-2fc03cc11c77
	github.com/opencontainers/runtime-spec v1.0.2-0.20190408193819-a1b50f621a48
	github.com/opentracing/opentracing-go v1.0.2
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.3.0
	github.com/uber/jaeger-client-go v2.15.0+incompatible
	github.com/vishvananda/netlink v1.0.1-0.20180723181557-2cbcf73e3dcd
	github.com/vishvananda/netns v0.0.0-20171111001504-be1fbeda1936
	golang.org/x/net v0.0.0-20190419010253-1f3472d942ba
	golang.org/x/sys v0.0.0-20200106162015-b016eb3dc98e
	google.golang.org/grpc v1.11.3
)

require (
	github.com/checkpoint-restore/go-criu v0.0.0-20181120144056-17b0214f6c48 // indirect
	github.com/cilium/ebpf v0.0.0-20200110133405-4032b1d8aae3 // indirect
	github.com/codahale/hdrhistogram v0.9.0 // indirect
	github.com/containerd/console v0.0.0-20180220200639-2748ece16665 // indirect
	github.com/coreos/go-systemd v0.0.0-20170731111925-d21964639418 // indirect
	github.com/cyphar/filepath-securejoin v0.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/go-units v0.3.2 // indirect
	github.com/godbus/dbus v4.1.0+incompatible // indirect
	github.com/golang/glog v1.2.2 // indirect
	github.com/golang/protobuf v0.0.0-20171113180720-1e59b77b52bf // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/mrunalp/fileutils v0.0.0-20171103030105-7d4729fb3618 // indirect
	github.com/opencontainers/selinux v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/seccomp/libseccomp-golang v0.9.1 // indirect
	github.com/syndtr/gocapability v0.0.0-20180916011248-d98352740cb2 // indirect
	github.com/uber-go/atomic v0.0.0-00010101000000-000000000000 // indirect
	github.com/uber/jaeger-lib v1.5.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20171123000638-7f0da29060c6 // indirect
)

replace github.com/uber-go/atomic => go.uber.org/atomic v1.11.0
