package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ShowLoadingScreen(app *tview.Application) {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBorder(true)

	infoTextView := tview.NewTextView().
		SetText("Loading...").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	flex.AddItem(infoTextView, 1, 1, false)

	go func() {
		if err := app.SetRoot(flex, true).Run(); err != nil {
			panic(err)
		}
	}()
}

func CreateResourceBox(title string) *tview.Flex {
	textView := tview.NewTextView().
		SetText(fmt.Sprintf("%s\n\n", title)).
		SetTextColor(tcell.Color104).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)

	flex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(textView, 0, 1, false)

	return flex
}

func RenderSystemInfo(updateCh chan<- *tview.Flex) {
	cpuUsage := GetCPUUsage()
	memoryUsage := GetMemoryUsage()
	runningProcesses := GetRunningProcess()
	physicalCores := GetPhysicalCPUCount()
	logicalCores := GetLogicalCPUCount()
	cpuModel := GetCPUInfo()

	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBorder(true)

	title := tview.NewTextView().
		SetText("System Monitor").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	cpuBox := CreateResourceBox(fmt.Sprintf("CPU : %.2f%%", cpuUsage))
	cpuModelBox := CreateResourceBox(fmt.Sprintf("CPU Model : %s", cpuModel))
	physicalCoresBox := CreateResourceBox(fmt.Sprintf("Physical CPU Cores : %d", physicalCores))
	logicalCoresBox := CreateResourceBox(fmt.Sprintf("Logical CPU Cores : %d", logicalCores))
	memoryBox := CreateResourceBox(fmt.Sprintf("Memory : %.2f%%", memoryUsage))
	processBox := CreateResourceBox(fmt.Sprintf("Processes :%s", runningProcesses))

	_ = title
	_ = cpuBox
	_ = memoryBox
	_ = processBox
	_ = physicalCoresBox
	_ = logicalCoresBox
	_ = cpuModelBox

	flex.AddItem(title, 1, 1, false).
		AddItem(cpuBox, 1, 1, false).
		AddItem(cpuModelBox, 1, 1, false).
		AddItem(physicalCoresBox, 1, 1, false).
		AddItem(logicalCoresBox, 1, 1, false).
		AddItem(memoryBox, 1, 1, false).
		AddItem(processBox, 0, 1, true)

	updateCh <- flex
}

func SetExitKeyHandler(app *tview.Application, updateCh chan<- *tview.Flex) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape || event.Rune() == 'q' {
			close(updateCh)
			app.Stop()

		}
		return event
	})
}
