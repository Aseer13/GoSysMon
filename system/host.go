package system

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/host"
)

type Host struct {
    Name string
    Uptime  uint64
    Arch string
    Os string
    kernel string
}

func CollectHost() {
    
    info, inferr := host.Info()
    if (inferr != nil) {
        fmt.Println("Failed to grab host info")
    }

    p := Host{info.Hostname, info.Uptime, info.PlatformFamily, info.KernelArch, info.KernelVersion}
    fmt.Println(p)
}

