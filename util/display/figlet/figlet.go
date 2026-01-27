package figlet

import (
	"fmt"
)

type Color int

// Foreground text colors.
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors.
const (
	FgHiBlack Color = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

func Logo() {
	logo := `
     ██╗██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ██╗ █████╗ ███╗   ██╗      ██╗  ██╗ █████╗ ███████╗       ██████╗██████╗ 
     ██║██║██╔══██╗████╗  ██║██╔════╝ ████╗  ██║██╔══██╗████╗  ██║      ██║ ██╔╝██╔══██╗██╔════╝      ██╔════╝██╔══██╗
     ██║██║███████║██╔██╗ ██║██║  ███╗██╔██╗ ██║███████║██╔██╗ ██║█████╗█████╔╝ ╚█████╔╝███████╗█████╗██║     ██║  ██║
██   ██║██║██╔══██║██║╚██╗██║██║   ██║██║╚██╗██║██╔══██║██║╚██╗██║╚════╝██╔═██╗ ██╔══██╗╚════██║╚════╝██║     ██║  ██║
╚█████╔╝██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚████║██║  ██║██║ ╚████║      ██║  ██╗╚█████╔╝███████║      ╚██████╗██████╔╝
 ╚════╝ ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝  ╚═══╝      ╚═╝  ╚═╝ ╚════╝ ╚══════╝       ╚═════╝╚═════╝
	`
	logo = ColorSize(logo, FgHiBlue)
	fmt.Println(logo)
}

// Colorize a string based on given color.

func ColorSize(s string, c Color) string {
	str := fmt.Sprintf("\033[%dm%s\033[0m", int(c), s)
	return str
}
