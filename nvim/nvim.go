package nvim

import (
	"themer/shell"
	"os/exec"
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

func (n *Nvim) SendCmd(args []string, errorMsg string) {
	shell.RunCmdOrPanic(append([]string{"nvim", "--server", n.Pipe}, args...), errorMsg)
}
