package command

import (
	"testing"
)

func TestRunCommand(t *testing.T) {
	// Arrange
	message := "Testing running a command"
	expectedOutput := message + "\n"

	// Act
	output, err := Run("echo", message)
	if err != nil {
		t.Errorf("ERROR: output was %v", err)
	}

	// Assert
	if output != expectedOutput {
		t.Errorf("Expected %q, got %q", expectedOutput, output)
	}
}
