package main

import (
	"fmt"
	"time"

	"fyne.io/systray"
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

	updateInstanceMenuItems()

	systray.AddSeparator()

	mTerminateAll := systray.AddMenuItem("Terminate all instances", "Shuts down all running WSL instances")
	mTerminateAll.SetIcon(icon.WSLClose)
	terminateAllOptions(mTerminateAll)

	systray.AddSeparator()
	addQuitOption()

	// Start checking for status changes
	go checkInstances(mTerminateAll)
}

func updateInstanceMenuItems() {
	WSLInstanceChecker.Update()

	currentInstances := make(map[string]struct{}, len(WSLInstanceChecker.Instances))
	for _, instance := range WSLInstanceChecker.Instances {
		currentInstances[instance.Name] = struct{}{}
		instName := instance.Name + " [" + string(instance.State) + "]"

		mInstance, ok := instanceMenuItems[instance.Name]
		if !ok {
			mInstance = systray.AddMenuItem(instName, instance.Name+" Instance")
			instanceMenuItems[instance.Name] = mInstance

			distroTerminateSubMenu(instance.Name, mInstance)
		}
		mInstance.SetTitle(instName)
		distroIcon := icon.GetByDistroName(instance.Name)
		mInstance.SetIcon(distroIcon)

		if instance.Running() {
			mInstance.Enable()
		} else {
			mInstance.Disable()
		}
	}

	for name, item := range instanceMenuItems {
		if _, ok := currentInstances[name]; !ok {
			item.Remove()
			delete(instanceMenuItems, name)
		}
	}
}

func distroTerminateSubMenu(instanceName string, mInstance *systray.MenuItem) {
	mTerminate := mInstance.AddSubMenuItem("Terminate", "Terminate "+instanceName)
	mTerminate.SetIcon(icon.WSLClose)
	go func() {
		for range mTerminate.ClickedCh {
			fmt.Println("Terminating", instanceName)
			err := wsl.Terminate(instanceName)
			if err != nil {
				fmt.Println("Error terminating", instanceName, ":", err)
			} else {
				fmt.Println(instanceName, "terminated successfully")
			}
		}
	}()
}

func checkInstances(items ...*systray.MenuItem) {
	for {
		updateInstanceMenuItems()
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
