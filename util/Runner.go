package util

import "os/exec"

type Runner interface {
	Run(string, ...string) ([]byte, error)
}

type RealRunner struct{}

func (real RealRunner) Run(cmd string, args ...string) ([]byte, error) {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	return out, err
}