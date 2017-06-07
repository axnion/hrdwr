package libs

import "os/exec"

type Runner interface {
	Run(string, ...string) ([]byte, error)
}

type RealRunner struct{}

func (RealRunner) Run(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).CombinedOutput()
}

/**
 * Takes a runner, a command string, and an arguments string. It runs the command using the runner
 * and the argument.
 */
func run(runner Runner, command string, arg string) ([]byte, error) {
	return runner.Run(command, arg)
}
func cmd(runner Runner, command string) ([]byte, error) {
	return runner.Run(command)
}
