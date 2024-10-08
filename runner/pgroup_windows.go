//go:build windows

package runner

import (
	"os/exec"
	"syscall"
)

func setPgid(cmd *exec.Cmd) {
	panic("Pgid not implemented on windows!")
}

func killGroup(cmd *exec.Cmd, _ syscall.Signal) error {
	// Windows doesn't have process groups, so just kill the process.
	return cmd.Process.Kill()
}
