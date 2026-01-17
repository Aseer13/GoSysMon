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
    Mhz [2]float64
    Temp int

}

func Grab() {
    p := Cpu{}
    info, _ := cpu.Info()
    temps, _ := sensors.SensorsTemperatures()

    if len(temps) == 0 {
        fmt.Println("No sensors found")
    }
   
    p.Model = info[0].ModelName  
    p.Cores, _ = cpu.Counts(false)
    p.Threads, _ =  cpu.Counts(true)
    getMhzRange(info, &p)    
}

func getMhzRange(info []cpu.InfoStat, pr *Cpu) {
    mhzMax, mhzMin := info[0].Mhz, info[0].Mhz
    for _, y := range info {
        if y.Mhz > mhzMax {
            mhzMax = y.Mhz
        } else if y.Mhz < mhzMin {
            mhzMin = y.Mhz
        }
    }
    
    if mhzMax == mhzMin {
        pr.Mhz[0] = mhzMax
        pr.Mhz[1] = mhzMin
    }
}