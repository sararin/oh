// Released under an MIT license. See LICENSE.

// +build !linux,!darwin,!dragonfly,!freebsd,!openbsd,!netbsd,!solaris,!windows

package task

import (
	"errors"
	. "github.com/michaelmacinnis/oh/pkg/cell"
	"os"
	"syscall"
)

var Platform string = "other"

func BecomeProcessGroupLeader() int {
	// TODO: Not sure what to do on non-Unix platforms.
	return 0
}

func ContinueProcess(pid int) {}

func GetHistoryFilePath() (string, error) {
	return "", errors.New("Not implemented")
}

func InitSignalHandling() {}

func JobControlSupported() bool {
	return false
}

func JoinProcess(proc *os.Process) int {
	status, err := proc.Wait()
	if err != nil {
		return -1
	}

	return status.Sys().(syscall.WaitStatus).ExitStatus()
}

func OSSpecificInit() {}

func ResetForegroundGroup(f *os.File) bool {
	return false
}

func SetForegroundGroup(group int) {}

func SuspendProcess(pid int) {}

func SysProcAttr(group int) *syscall.SysProcAttr {
	return nil
}

func TerminateProcess(pid int) {}

func evaluate(c Cell, file string, line int, problem string) Cell {
	task0.Eval <- Message{Cmd: c, File: file, Line: line, Problem: problem}
	return <-task0.Done
}
