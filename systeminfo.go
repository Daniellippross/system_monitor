package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCPUUsage() float64 {
	var percentages []float64
	var err error

	percentages, err = cpu.Percent(time.Second, false)

	if err != nil {
		fmt.Println(" Error Fetching CPU Percentage ", err)
		return -1
	}

	if len(percentages) > 0 {
		return percentages[0]
	}

	return -1
}

func GetPhysicalCPUCount() int {
	var cores int
	var err error

	cores, err = cpu.Counts(false)
	if err != nil {
		fmt.Println("Error with core count")
		return -1
	}

	return cores
}

func GetLogicalCPUCount() int {
	var cores int
	var err error

	cores, err = cpu.Counts(true)
	if err != nil {
		fmt.Println("Error with core count")
		return -1
	}

	return cores
}

func GetCPUInfo() string {
	infos, err := cpu.Info()
	if err != nil {
		return fmt.Sprintln("Error getting CPU model", err)
	}

	return infos[0].ModelName
}

func GetMemoryUsage() float64 {
	var memory *mem.VirtualMemoryStat
	var err error

	memory, err = mem.VirtualMemory()

	if err != nil || memory == nil {
		fmt.Printf("Error Getting Memory info: %v", err)
		return -1
	}
	return memory.UsedPercent
}

func GetRunningProcess() string {
	var processes []*process.Process
	var err error

	processes, err = process.Processes()

	if err != nil {
		return fmt.Sprintf("Error getting process info: %v", err)
	}

	var processInfo string
	processCount := 0

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			name = "N/A"
		}

		processInfo += fmt.Sprintf("\n PID: %d, Name: %s", p.Pid, name)

		processCount++

		if processCount >= 10 {
			break
		}
	}

	return processInfo
}
