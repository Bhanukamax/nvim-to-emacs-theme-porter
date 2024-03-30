package shell

import (
	"os/exec"
	"fmt"
)

func RunCmdOrPanic(args []string, errorMsg string) {
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Bmax: Error ", errorMsg, err))
	}
}

func SafeDeleteFile(fileName string) {
	RunCmdOrPanic([]string{"/usr/bin/env", "touch", fileName}, "Bmax: Error removing pipe")
	RunCmdOrPanic([]string{"/usr/bin/env", "rm", fileName}, "Bmax: Error removing pipe")
}
