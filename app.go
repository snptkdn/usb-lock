package main

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	// "os/exec"
	"regexp"
	"strings"
)

type UsbDeviceInfo struct {
	Model string
	Index int
	Size  string
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetUsbDevices() []UsbDeviceInfo {
	output := `Model                        Index  Size
SanDisk Ultra USB 3.0        2      128GB
WD My Passport 1234 USB Device 3    1TB
`

	return parseUsbDeviceInfo(string(output))
}

func parseUsbDeviceInfo(output string) []UsbDeviceInfo {
	var devices []UsbDeviceInfo
	scanner := bufio.NewScanner(strings.NewReader(output))
  regex := regexp.MustCompile(`^(.+)\s+(\d+)\s+([\d\w]+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Model") || line == "" {
			continue
		}

		matches := regex.FindStringSubmatch(line)
		if matches != nil && len(matches) > 3 {
			model := matches[1]
			index, _ := strconv.Atoi(matches[2])
			size := matches[3]

			devices = append(devices, UsbDeviceInfo{
				Model: model,
				Index: index,
				Size:  size,
			})
		}
	}

	return devices
}
