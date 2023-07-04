// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || amd64p32 || arm || arm64 || loong64 || mips64le || mips64p32le || mipsle || ppc64le || riscv64

package tracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type networkEventT struct {
	Netns     uint32
	_         [4]byte
	Timestamp uint64
	MountNsId uint64
	Pid       uint32
	Tid       uint32
	Uid       uint32
	Gid       uint32
	Task      [16]uint8
	PktType   uint32
	Ip        uint32
	Proto     uint16
	Port      uint16
	_         [4]byte
}

type networkSocketsKey struct {
	Netns  uint32
	Family uint16
	Proto  uint16
	Port   uint16
	_      [2]byte
}

type networkSocketsValue struct {
	Mntns             uint64
	PidTgid           uint64
	UidGid            uint64
	Task              [16]int8
	Sock              uint64
	DeletionTimestamp uint64
	Ipv6only          int8
	_                 [7]byte
}

// loadNetwork returns the embedded CollectionSpec for network.
func loadNetwork() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_NetworkBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load network: %w", err)
	}

	return spec, err
}

// loadNetworkObjects loads network and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*networkObjects
//	*networkPrograms
//	*networkMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadNetworkObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadNetwork()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// networkSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type networkSpecs struct {
	networkProgramSpecs
	networkMapSpecs
}

// networkSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type networkProgramSpecs struct {
	IgTraceNet *ebpf.ProgramSpec `ebpf:"ig_trace_net"`
}

// networkMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type networkMapSpecs struct {
	Events  *ebpf.MapSpec `ebpf:"events"`
	Sockets *ebpf.MapSpec `ebpf:"sockets"`
}

// networkObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadNetworkObjects or ebpf.CollectionSpec.LoadAndAssign.
type networkObjects struct {
	networkPrograms
	networkMaps
}

func (o *networkObjects) Close() error {
	return _NetworkClose(
		&o.networkPrograms,
		&o.networkMaps,
	)
}

// networkMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadNetworkObjects or ebpf.CollectionSpec.LoadAndAssign.
type networkMaps struct {
	Events  *ebpf.Map `ebpf:"events"`
	Sockets *ebpf.Map `ebpf:"sockets"`
}

func (m *networkMaps) Close() error {
	return _NetworkClose(
		m.Events,
		m.Sockets,
	)
}

// networkPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadNetworkObjects or ebpf.CollectionSpec.LoadAndAssign.
type networkPrograms struct {
	IgTraceNet *ebpf.Program `ebpf:"ig_trace_net"`
}

func (p *networkPrograms) Close() error {
	return _NetworkClose(
		p.IgTraceNet,
	)
}

func _NetworkClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed network_bpfel.o
var _NetworkBytes []byte
