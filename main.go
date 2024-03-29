package main

import (
	"os"
	"fmt"
	"os/exec"
	"time"
	"themer/lexer"
)

func checkAndPanic(err error, errorMsg string) {
	if err != nil {
		panic(fmt.Sprintf("Bmax: Error ", errorMsg, err))
	}
}
func runCmd(errorMsg string, arg ...string) {
	cmd := exec.Command(arg[0], arg[1:]...)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Bmax: Error ", errorMsg, err))
	}
}

func runNvim() {
	runCmd("Bmax: Error removing pipe", "/usr/bin/env", "touch", "/tmp/bmax-nvim.pipe" )
	runCmd("Bmax: Error removing pipe", "/usr/bin/env", "rm", "/tmp/bmax-nvim.pipe" )
	runCmd("Bmax: Error running the Command", "/usr/bin/env", "nvim", "--headless", "--listen", "/tmp/bmax-nvim.pipe" )
}

func exportTheme() {
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
		time.Sleep(time.Second/2)
		fmt.Println("exporting theme to theme.vim")
		exportTheme()
		fmt.Println("done")
	}
	colorMap := makeColorMap()
	fmt.Println("Colol Map", colorMap)
}

	//	colorMap := makeColorMap()
	//	fmt.Println(colorMap)

type Color struct {
	fg string
	bg string
}

func newColor(fg string, bg string) *Color {
	return &Color{fg, bg}
	//	c := Color{fg, bg}
	//	return &c
}

type ColorMap map[string]Color

func makeColorMap() ColorMap{
	colorMap := make(map[string]Color)
	data, err := os.ReadFile("./theme.vim")
	checkAndPanic(err, "reading file")
	fmt.Println(string(data))

	lexer.New()
	return colorMap
}
