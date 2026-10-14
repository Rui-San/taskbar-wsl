package icon

import "strings"

var distroIcons = map[string][]byte{
	"ubuntu": Ubuntu,
	"fedora": Fedora,
}

func GetByDistroName(name string) []byte {
	name = strings.ToLower(name)

	for distro, icon := range distroIcons {
		if strings.Contains(name, distro) {
			return icon
		}
	}
	return Tux
}
