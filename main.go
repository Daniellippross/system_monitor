package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	ShowLoadingScreen(app)

	updateCh := make(chan *tview.Flex, 1)

	defer close(updateCh)

	go func() {
		for {
			RenderSystemInfo(updateCh)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case updatedFlex, ok := <-updateCh:
				if !ok {
					return
				}
				app.QueueUpdateDraw(func() {
					app.SetRoot(updatedFlex, true)
				})
			}
		}
	}()

	SetExitKeyHandler(app, updateCh)

	select {}
}

func printSystemInfo() {
	cpuUsage := GetCPUUsage()
	memoryUsage := GetMemoryUsage()
	runningProcesses := GetRunningProcess()
	physicalCores := GetPhysicalCPUCount()
	logicalCores := GetLogicalCPUCount()

	fmt.Println("CPU Percentage    :", cpuUsage)
	fmt.Println("Memory Percentage :", memoryUsage)
	fmt.Println("Running Processes :", runningProcesses)
	fmt.Println("Physical Cores :", physicalCores)
	fmt.Println("Logical Cores :", logicalCores)
}
