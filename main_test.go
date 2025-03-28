package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCPUUsage(t *testing.T) {
	cpuUsage := GetCPUUsage()
	assert.GreaterOrEqual(t, cpuUsage, 0.0, "CPU usage should be greater than equal to 0")
}

func TestGetPhysicalCPUCount(t *testing.T) {
	physicalCpuCount := GetPhysicalCPUCount()
	assert.GreaterOrEqual(t, physicalCpuCount, 0, "Physical CPU core count should be grater than 0")
}

func TestGetLogicalCPUCount(t *testing.T) {
	logicalCpuCount := GetPhysicalCPUCount()
	assert.GreaterOrEqual(t, logicalCpuCount, 0, "Logical CPU core count should be grater than 0")
}

func TestGetCPUInfo(t *testing.T) {
	cpuInfo := GetCPUInfo()
	assert.Greater(t, len(cpuInfo), 0, "There should be a model for the cpu")
}

func TestGetMemoryUsage(t *testing.T) {
	cpuUsage := GetMemoryUsage()
	assert.GreaterOrEqual(t, cpuUsage, 0.0, "Memory usage should be greater than equal to 0")
}

func TestRunningProcess(t *testing.T) {
	processes := GetRunningProcess()
	assert.Greater(t, len(processes), 0, "There should be atleast one running process")
}

func TestPrintSystemInfo(t *testing.T) {
	// Redirect standard output for testing
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function to be tested
	printSystemInfo()

	// Reset standard output after the test
	w.Close()
	os.Stdout = old

	// Capture the output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Define expected output

	// Compare actual and expected output
	got := buf.String()
	println(" got " + got)

	if strings.Contains(got, "CPU Percentage    : -1") {
		t.Errorf("CPU Percentage information is not present. Expected:\n%s\nActual:\n%s", "CPU Percentage greater than or equal to 0", got)
	}
	if strings.Contains(got, "Memory Percentage : -1") {
		t.Errorf("Memory Percentage information is not present. Expected:\n%s\nActual:\n%s", "Memory Percentage greater than or equal to 0", got)
	}

}

func TestPrintSystemInfoInterval(t *testing.T) {
	// Redirect standard output for testing
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Capture the output
	var buf bytes.Buffer
	go func() {
		// Call the function to be tested periodically
		for i := 0; i < 2; i++ {
			printSystemInfo()
			time.Sleep(2 * time.Second)
		}

		// Reset standard output after the test
		w.Close()
		os.Stdout = old
	}()

	// Wait for the goroutine to finish
	time.Sleep(6 * time.Second)

	// Copy the captured output
	io.Copy(&buf, r)

	// Compare actual and expected output
	got := buf.String()
	println(" Output :\n " + got)

	count := strings.Count(got, "CPU Percentage    :")

	if count < 2 {
		t.Errorf("The System Resource information is not being printed continously. \nActual:\n%s", got)
	}

	if strings.Contains(got, "CPU Percentage    : -1") {
		t.Errorf("CPU Percentage information is not present. Expected:\n%s\nActual:\n%s", "CPU Percentage greater than or equal to 0", got)
	}
	if strings.Contains(got, "Memory Percentage : -1") {
		t.Errorf("Memory Percentage information is not present. Expected:\n%s\nActual:\n%s", "Memory Percentage greater than or equal to 0", got)
	}
}
