//go:build windows

package icon

import (
	_ "embed"
)

//go:embed "wsl_icon_default.ico"
var DefaultData []byte

//go:embed "wsl_icon_paused.ico"
var PausedData []byte

//go:embed "wsl_icon_running.ico"
var RunningData []byte
