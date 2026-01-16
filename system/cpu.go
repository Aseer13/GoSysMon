package system

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
)


type Cpu struct {
    Cores int32
}


func GetCpu() {

    fmt.Println(cpu.Info())
}

