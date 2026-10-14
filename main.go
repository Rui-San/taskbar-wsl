package main

import (
	"fmt"
	"time"

	"fyne.io/systray"
	"github.com/Rui-San/taskbar-wsl/command"
	"github.com/Rui-San/taskbar-wsl/icon"
	"github.com/Rui-San/taskbar-wsl/wsl"
)

var WSLInstanceChecker = wsl.NewInstanceChecker()

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

	mShutdownAll := systray.AddMenuItem("Shutdown all instances", "Shuts down all running WSL instances")
	mShutdownAll.SetIcon(icon.WSLClose)
	shutdownAllOptions(mShutdownAll)

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
		time.Sleep(5 * time.Second)
	}
}

func shutdownAllOptions(mShutdownAll *systray.MenuItem) {
	go func() {
		for range mShutdownAll.ClickedCh {
			fmt.Println("Starting shutdown of all WSL instances...")
			err := shutdownAllWSLInstances()
			if err != nil {
				fmt.Println("Error shutting down WSL instances:", err)
			} else {
				fmt.Println("All WSL instances shut down successfully")
			}
		}
	}()
}

func shutdownAllWSLInstances() error {
	_, err := command.Run("wsl", "--shutdown")
	return err
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
