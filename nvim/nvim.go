package nvim

type Nvim struct {
	Pipe string
}

func StartServer(pipe string) *Nvim {

	n := &Nvim{pipe }
	shell.SafeDeleteFile("./theme.vim")
	shell.RunCmdOrPanic([]string{"/usr/bin/env", "nvim", "--headless", "--listen", n.Pipe}, "Bmax: Error running the Commandtt")

	return &n
}

func (n *Nvim) SendCmd(args []string, errorMsg string) {
	shell.RunCmdOrPanic({"nvim", "--server", n.Pipe, ...args}, errorMsg)
}
