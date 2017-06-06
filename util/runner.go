package util

import "os/exec"

type Runner interface {
	Run(string, ...string) ([]byte, error)
}

type RealRunner struct{}

func (RealRunner) Run(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).CombinedOutput()
}