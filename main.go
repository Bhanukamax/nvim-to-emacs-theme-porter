package main

import (
	"fmt"
	"os/exec"
	"time"
)

func runNvim() {
	cmd := exec.Command("/usr/bin/env", "nvim", "--headless", "--listen", "/tmp/bmax-nvim.pipe" )
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

func main() {
	go runNvim()
	fmt.Println("starting neovim")
	time.Sleep(time.Second)
	fmt.Println("exporting theme to theme.txt")
	exportTheme()
	fmt.Println("done")
}
