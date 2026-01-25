package system

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
    "github.com/shirou/gopsutil/v4/sensors"
)

type Cpu struct {
    Model string
    Cores int
    Threads int
    Mhz float64
    Temp float64

}

func CollectCpu() {
    
    info, inferr := cpu.Info()
    if (inferr != nil) {
        fmt.Println("Failed to get info")
    }

    temps, temperr := sensors.SensorsTemperatures()
    if (temperr != nil) {
        fmt.Println("Failed to get sensors")
    }

    c := Cpu{}
    c.Model = info[0].ModelName  
    c.Cores, _ = cpu.Counts(false)
    c.Threads, _ =  cpu.Counts(true)
    c.Mhz = info[0].Mhz

    if (len(temps) > 0) {
        c.Temp = temps[0].Temperature
    }
}

