module github.com/hfyeh/agent

go 1.21

require (
	github.com/docker/docker v27.1.0+incompatible
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645
	github.com/hashicorp/yamux v0.1.1
	github.com/mdlayher/vsock v0.0.0
	github.com/opencontainers/runc v0.0.0
	github.com/opencontainers/runtime-spec v1.2.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.9.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/vishvananda/netlink v1.1.0
	github.com/vishvananda/netns v0.0.4
	golang.org/x/net v0.27.0
	golang.org/x/sys v0.22.0
	google.golang.org/grpc v1.65.0
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/checkpoint-restore/go-criu v4.0.0+incompatible // indirect
	github.com/cilium/ebpf v0.7.0 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/coreos/go-systemd v15 // indirect
	github.com/cyphar/filepath-securejoin v0.2.4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/godbus/dbus v4.1.0+incompatible // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/mrunalp/fileutils v0.5.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/opencontainers/selinux v1.10.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/seccomp/libseccomp-golang v0.9.2-0.20220502022130-f33da4d89646 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/mdlayher/vsock v0.0.0 => github.com/mdlayher/vsock v0.0.0-20190429153235-7b7533a7ca4e
	github.com/opencontainers/runc v0.0.0 => github.com/opencontainers/runc v1.0.0-rc9.0.20200122160610-2fc03cc11c77
)
