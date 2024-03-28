package main

import (
	"fmt"
	"os/exec"
	"time"
	//	"reflect"
)

func runNvim() {
	cmd := exec.Command("/usr/bin/env", "touch", "/tmp/bmax-nvim.pipe" )
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Bmax: Error removing pipe", err))
	}
	cmd = exec.Command("/usr/bin/env", "rm", "/tmp/bmax-nvim.pipe" )
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Bmax: Error removing pipe", err))
	}
	cmd = exec.Command("/usr/bin/env", "nvim", "--headless", "--listen", "/tmp/bmax-nvim.pipe" )
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Bmax: Error running the Command", err))
	}
}

func exportTheme() {
	cmd := exec.Command("./export.sh")
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprintf("Bmax: Error exporting theme", err))
	}
}

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
	return colorMap
}

func main() {
	//	colorMap := makeColorMap()
	//	fmt.Println(colorMap)

	go runNvim()
	fmt.Println("starting neovim")
	time.Sleep(time.Second/2)
	fmt.Println("exporting theme to theme.vim")
	exportTheme()
	fmt.Println("done")
}
