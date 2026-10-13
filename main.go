package main

import (
	"fmt"
	"time"

	"fyne.io/systray"
	"github.com/Rui-San/taskbar-wsl/icon"
)

func main() {
	systray.Run(onReady, onExit)
}

func onExit() {
	now := time.Now()
	fmt.Println("Closed at", now.String())
}

func onReady() {
	systray.SetIcon(icon.DefaultData)
	systray.SetTitle("Taskbar WSL")
	systray.SetTooltip("Taskbar WSL")

	addQuitOption()
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
