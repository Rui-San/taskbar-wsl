// Basically a cmd wrapper for easier testing
package command

import "os/exec"

func Run(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
