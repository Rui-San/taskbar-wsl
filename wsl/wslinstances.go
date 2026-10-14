package wsl

import (
	"fmt"
	"strings"
	"time"

	"github.com/Rui-San/taskbar-wsl/command"
)

const (
	instanceStateUnknown instanceState = "Unknown"
	instanceStateRunning instanceState = "Running"
	instanceStateStopped instanceState = "Stopped"
)

type instanceState string

// WSL Instance

type instance struct {
	Name    string
	State   instanceState
	Version string
}

func (i *instance) Running() bool {
	return i.State == instanceStateRunning
}

// WSL Instance Holder (For operations on multiple instances)

type instanceHolder struct {
	Instances []instance
}

func NewInstanceChecker() *instanceHolder {
	return &instanceHolder{
		Instances: []instance{},
	}
}

func (ih *instanceHolder) Update() {
	ih.Instances = []instance{} // reset

	out, err := command.Run("wsl", "-l", "-v")
	if err != nil {
		fmt.Println("Error trying to run `wsl -l -v`. WSL might not be installed: ", err.Error())
		return
	}

	for i, outLine := range strings.Split(out, "\n") {
		if i == 0 || len(outLine) < 3 { // skip header since it's just "NAME STATE VERSION"
			continue
		}
		instance, err := parseOutputLine(outLine)
		if err != nil {
			fmt.Println("Error parsing output line: ", err.Error())
			continue
		}
		ih.Instances = append(ih.Instances, instance)
	}
	now := time.Now()
	fmt.Printf("[%s] Finished updating instances: %s\n", now.Format("2006-01-02 15:04:05"), ih.Instances)
}

func parseOutputLine(line string) (instance, error) {
	fields := strings.Fields(strings.ReplaceAll(line, "\x00", ""))
	if len(fields) < 3 {
		return instance{}, fmt.Errorf("At least 3 fields are expected")
	}

	// The default distro is prefixed with an *
	if fields[0] == "*" {
		fields = fields[1:]
	}

	if len(fields) < 3 {
		return instance{}, fmt.Errorf("Expected distro name, state, and version")
	}

	newInstance := instance{
		Name:    fields[0],
		Version: fields[2],
	}
	switch fields[1] {
	case "Running":
		newInstance.State = instanceStateRunning
	case "Stopped":
		newInstance.State = instanceStateStopped
	default:
		newInstance.State = instanceStateUnknown
	}

	return newInstance, nil
}

func (ih *instanceHolder) RunningCount() int {
	n := 0
	for _, instance := range ih.Instances {
		if instance.State == instanceStateRunning {
			n++
		}
	}
	return n
}
