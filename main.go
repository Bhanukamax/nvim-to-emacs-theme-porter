package main

import (
	"fmt"
	"os"
	"os/exec"
	"themer/lexer"
	"time"
	"strings"
)

func checkAndPanic(err error, errorMsg string) {
	if err != nil {
		panic(fmt.Sprintf("Bmax: Error ", errorMsg, err))
	}
}
func runCmdOrPanic(args []string, errorMsg string) {
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Bmax: Error ", errorMsg, err))
	}
}

func safeDeleteFile(fileName string) {
	runCmdOrPanic([]string{"/usr/bin/env", "touch", fileName}, "Bmax: Error removing pipe")
	runCmdOrPanic([]string{"/usr/bin/env", "rm", fileName}, "Bmax: Error removing pipe")
}

func runNvim() {
	safeDeleteFile("./theme.vim")
	runCmdOrPanic([]string{"/usr/bin/env", "nvim", "--headless", "--listen", "/tmp/bmax-nvim.pipe"}, "Bmax: Error running the Command")
}

func exportTheme() {
	safeDeleteFile("./theme.vim")
	cmd := exec.Command("./export.sh")
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprintf("Bmax: Error exporting theme", err))
	}
}

var shouldExport bool

func main() {
	shouldExport = false
	if shouldExport {
		go runNvim()
		fmt.Println("starting neovim")
		time.Sleep(time.Second / 2)
		fmt.Println("exporting theme to theme.vim")
		exportTheme()
		fmt.Println("done")
	}

	colorMap := makeColorMap()
	for	key, color := range colorMap {
		fmt.Println("key: ", key, "color", color)
	}
}

//	colorMap := makeColorMap()
//	fmt.Println(colorMap)

type Color struct {
	Fg string
	Bg string
}

func newColor(fg string, bg string) *Color {
	return &Color{fg, bg}
	//	c := Color{fg, bg}
	//	return &ce
}

type ColorMap map[string]Color

func parseColor(input string) Color {
	color := Color{}

	parts := strings.Fields(input)
	for _, part := range parts {
		if strings.HasPrefix(part, "guifg=") {
			color.Fg = strings.TrimPrefix(part, "guifg=")
		} else if strings.HasPrefix("guibg=", part) {
			color.Bg = strings.TrimPrefix(part, "guibg=")
		}
	}

	return color
}

func makeColorMap() ColorMap {
	colorMap := make(map[string]Color)
	data, err := os.ReadFile("./theme.vim")
	checkAndPanic(err, "reading file")
	//	fmt.Println(string(data))
	lines := strings.Split(string(data), "\n")
	for i, colorLine := range(lines) {
		fmt.Println(i , lines[i])
		parts := strings.Fields(colorLine)
		key := parts[1]
		colorMap[key] = parseColor(colorLine)
	}

	lexer.New("foo")
	return colorMap
}
