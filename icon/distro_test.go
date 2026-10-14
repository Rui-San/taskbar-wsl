package icon

import "testing"

func TestGetByDistroName_ReturnsExpectedIcon(t *testing.T) {
	// Arrange
	expectedIcon := []byte{0x00}
	distroName := "testing"
	distroIcons[distroName] = expectedIcon

	// Act
	icon := GetByDistroName("TeStiNG-distro-test")

	// Assert
	if icon[0] != expectedIcon[0] {
		t.Errorf("Expected icon for %s, got %v", distroName, icon)
	}
}
