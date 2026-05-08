package shared

import (
	"fmt"
	"runtime"
	"strings"
)

// ANSI color constants
const (
	Blue  = "\033[34m"
	Cyan  = "\033[36m"
	Reset = "\033[0m"
	Bold  = "\033[1m"
)

// PrintHeader displays the ASCII art and version info
func PrintHeader() {
	goVersion := strings.TrimPrefix(runtime.Version(), "go")

	banner := `
  _____                 _     _ ____  _
 |  __ \               (_)   | |  _ \(_)
 | |__) |__ _ _ __  _  _  __| | |_) |_ _ __   __ _ _ __ _   _
 |  _  // _` + "`" + ` | '_ \| | | |/ _` + "`" + ` |  _ <| | '_ \ / _` + "`" + ` | '__| | | |
 | | \ \ (_| | |_) | | | | (_| | |_) | | | | | (_| | |  | |_| |
 |_|  \_\__,_| .__/  |_|_|\__,_|____/|_|_| |_|\__,_|_|   \__, |
             | |                                          __/ |
             |_|                                         |___/
`
	fmt.Printf("%s%s%s", Blue, banner, Reset)
	fmt.Printf("%s%s Running on Go %s %s\n", Bold, Cyan, goVersion, Reset)
	fmt.Println("--------------------------------------------------")
}
