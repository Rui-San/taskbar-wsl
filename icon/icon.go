package icon

import (
	_ "embed"
)

//go:embed "assets/bullshit.ico"
var TBWSL []byte

//go:embed "assets/linux_default.ico"
var Tux []byte

//go:embed "assets/ubuntu.ico"
var Ubuntu []byte

//go:embed "assets/fedora.ico"
var Fedora []byte

//go:embed "assets/wsl_icon_default.ico"
var WSL []byte

//go:embed "assets/wsl_icon_paused.ico"
var WSLPaused []byte

//go:embed "assets/wsl_icon_running.ico"
var WSLRunning []byte

//go:embed "assets/wsl_icon_close.ico"
var WSLClose []byte
