package ezchip-whats

import (
	"github.com/powersj/whatsthis/pkg/cloud"
	"github.com/powersj/whatsthis/pkg/container"
	"github.com/powersj/whatsthis/pkg/cpu"
	"github.com/powersj/whatsthis/pkg/distro"
	"github.com/powersj/whatsthis/pkg/memory"
	"github.com/powersj/whatsthis/pkg/network"
	"github.com/powersj/whatsthis/pkg/platform"
	"github.com/powersj/whatsthis/pkg/storage"
	"github.com/powersj/whatsthis/pkg/virt"
)

// CloudProbe alias for cloud.Probe.
type CloudProbe = cloud.Probe

// ContainerProbe alias for container.Probe.
type ContainerProbe = container.Probe

// CPUProbe alias for cpu.Probe.
type CPUProbe = cpu.Probe

// DistroProbe alias for distro.Probe.
type DistroProbe = distro.Probe

// MemoryProbe alias for memory.Probe.
type MemoryProbe = memory.Probe

// NetworkProbe alias for network.Probe.
type NetworkProbe = network.Probe

// StorageProbe alias for storage.Probe.
type StorageProbe = storage.Probe

// PlatformProbe alias for platform.Probe.
type PlatformProbe = platform.Probe

// VirtProbe alias for virt.Probe.
type VirtProbe = virt.Probe

// Cloud alias for new struct.
var Cloud = cloud.New

// Container alias for new struct.
var Container = container.New

// CPU alias for new struct.
var CPU = cpu.New

// Distro alias for new struct.
var Distro = distro.New

// Memory alias for new struct.
var Memory = memory.New

// Network alias for new struct.
var Network = network.New

// Storage alias for new struct.
var Storage = storage.New

// Platform alias for new struct.
var Platform = platform.New

// Virt alias for new struct.
var Virt = virt.New
