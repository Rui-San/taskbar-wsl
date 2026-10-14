package wsl

import "testing"

func TestParseOutputLine_WithDefaultDistro_ReturnsExpectedInstance(t *testing.T) {
	// Arrange
	line := "* Ubuntu Running 2"
	expected := instance{
		Name:    "Ubuntu",
		State:   instanceStateRunning,
		Version: "2",
	}

	// Act
	out, err := parseOutputLine(line)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	// Assert
	if out != expected {
		t.Fatalf("expected: %+v, got: %+v", expected, out)
	}
}

func TestParseOutputLine_WithNonDefaultDistro_ReturnsExpectedInstance(t *testing.T) {
	// Arrange
	line := "* Fedora Stopped 2"
	expected := instance{
		Name:    "Fedora",
		State:   instanceStateStopped,
		Version: "2",
	}

	// Act
	out, err := parseOutputLine(line)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	// Assert
	if out != expected {
		t.Fatalf("expected: %+v, got: %+v", expected, out)
	}
}

func TestRunningCount_ReturnsCorrectNumberOfRunningInstances(t *testing.T) {
	// Arrange
	ih := instanceHolder{
		Instances: []instance{
			{Name: "Ubuntu", State: instanceStateRunning, Version: "2"},
			{Name: "Fedora", State: instanceStateStopped, Version: "2"},
			{Name: "Debian", State: instanceStateRunning, Version: "2"},
			{Name: "Arch", State: instanceStateRunning, Version: "2"},
		},
	}
	expectedCount := 3

	// Act
	count := ih.RunningCount()

	// Assert
	if count != expectedCount {
		t.Fatalf("expected: %d, got: %d", expectedCount, count)
	}
}
