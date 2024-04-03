package nvim

import (
	"fmt"
	"os/exec"
	"themer/shell"
)

type Nvim struct {
	Pipe string
}

func New (pipe string) *Nvim{
	return &Nvim {
		Pipe: pipe,
	}
}

func (n *Nvim) StartServer() error {

	shell.SafeDeleteFile("./theme.vim")
	cmd :=	exec.Command("/usr/bin/env", "nvim", "--headless", "--listen", "/tmp/bmax-nvim.pipe")
	return cmd.Run()
}

func (n *Nvim) SendCmd(args []string) error {
	fmt.Println("/usr/bin/env", append([]string{"nvim", "--server", n.Pipe, "--remote-send", "'"}, args...))
	cmd := exec.Command("/usr/bin/env", append([]string{"nvim", "--server", n.Pipe, "--remote-send"}, args...)...)
	return cmd.Run()
}

func ExportTheme() {
	shell.SafeDeleteFile("./theme.vim")
	cmd := exec.Command("./export.sh")
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprintf("Bmax: Error exporting theme", err))
	}
}
