package main

import (
	"fmt"
	"os/exec"
	"time"
)

func runNvim() {
	//	cmd := exec.Command("/home/bmax/.local/bin/nvim --headless --listen /tmp/bmax-nvim.pipe")
	cmd := exec.Command("/home/bmax/.local/bin/nvim", "--headless", "--listen", "/tmp/bmax-nvim.pipe" )

	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Bmax: Error running the Command", err))
	}
}

func exportTheme() {
	//	cmd := exec.Command("/home/bmax/.local/bin/nvim --headless --listen /tmp/bmax-nvim.pipe")
	cmd := exec.Command("/home/bmax/bmax/nvim-to-emacs-theme/export.sh")

	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprintf("Bmax: Error exporting theme", err))
	}
}

func main() {
	go runNvim()
	fmt.Println("sleeping")
	time.Sleep(1 * time.Second)
	exportTheme()
	fmt.Println("done")
}
