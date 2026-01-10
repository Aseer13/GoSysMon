package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu"
)

func main() {
	v, _ := cpu.Info()
	fmt.Println(v)
}