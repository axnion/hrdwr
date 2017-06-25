package lib

import "os/exec"

var runner Runner = RealRunner{} // Global variable used to execute shell commands

/**
 * The runner interface is used by RealRunner and is there so RealRunner can be mocked when testing.
 */
type Runner interface {
	run(string, ...string) ([]byte, error)
}

/**
 * RealRunner is the object used when executing commands, adhere to the Runner interface.
 */
type RealRunner struct{}

/**
 * The run method of RealRunner. Calls the exec.Command method with the command string and any arguments.
 */
func (RealRunner) run(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).CombinedOutput()
}
