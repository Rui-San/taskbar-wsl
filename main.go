package main

import (
	"fmt"
	"time"

	"fyne.io/systray"
	"github.com/Rui-San/taskbar-wsl/command"
	"github.com/Rui-San/taskbar-wsl/icon"
	"github.com/Rui-San/taskbar-wsl/wsl"
)

const (
	REFRESH_INTERVAL = 2 * time.Second
)

var (
	WSLInstanceChecker = wsl.NewInstanceChecker()
	instanceMenuItems  = make(map[string]*systray.MenuItem)
)

func main() {
	systray.Run(onReady, onExit)
}

func onExit() {
	now := time.Now()
	fmt.Println("Closed at", now.String())
}

func onReady() {
	systray.SetIcon(icon.TBWSL)
	systray.SetTitle("Taskbar WSL")
	systray.SetTooltip("Taskbar WSL")

	addInstancesMenuItems()

	systray.AddSeparator()

	mTerminateAll := systray.AddMenuItem("Terminate all instances", "Shuts down all running WSL instances")
	mTerminateAll.SetIcon(icon.WSLClose)
	terminateAllOptions(mTerminateAll)

	systray.AddSeparator()
	addQuitOption()

	// Start checking for status changes
	go checkInstances(mShutdownAll)
}

// Need to make this dynamic and add submenus
func addInstancesMenuItems() {
	WSLInstanceChecker.Update()
	for _, instance := range WSLInstanceChecker.Instances {
		instName := instance.Name + " [" + string(instance.State) + "]"
		mInstance := systray.AddMenuItem(instName, "WSL instance: "+instance.Name)
		distroIcon := icon.GetByDistroName(instance.Name)
		mInstance.SetIcon(distroIcon)
		if !instance.Running() {
			mInstance.Disable()
		}
	}
}

func checkInstances(items ...*systray.MenuItem) {
	for {
		WSLInstanceChecker.Update()
		if WSLInstanceChecker.RunningCount() > 0 {
			systray.SetIcon(icon.WSLRunning)
			for _, item := range items {
				item.Enable()
			}
		} else {
			systray.SetIcon(icon.WSLPaused)
			for _, item := range items {
				item.Remove()
			}
		}
		time.Sleep(REFRESH_INTERVAL)
	}
}

func terminateAllOptions(mTerminateAll *systray.MenuItem) {
	go func() {
		for range mTerminateAll.ClickedCh {
			fmt.Println("Starting termination of all WSL instances...")
			err := wsl.TerminateAll()
			if err != nil {
				fmt.Println("Error terminating WSL instances:", err)
			} else {
				fmt.Println("All WSL instances terminated successfully")
			}
		}
	}()
}

func addQuitOption() {
	mQuit := systray.AddMenuItem("Close", "Closes WSL Taskbar")
	go func() {
		for range mQuit.ClickedCh {
			systray.Quit()
			fmt.Println("Closing...")
		}
	}()
}
