package lib

import "os/exec"

type Runner interface {
	run(string, ...string) ([]byte, error)
}

type RealRunner struct{}

func (RealRunner) run(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).CombinedOutput()
}
