package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	getcpu()
	getram()
}

func getcpu() {
	v, _ := cpu.Info()
	
	for x, y := range v {
		fmt.Printf("CODE #%d: %d\n", x, int(y.Mhz))
	}
}

func getram() {
	v, _ := mem.SwapMemory()
	
	fmt.Printf("RAM: %d\n", v.Total)
}