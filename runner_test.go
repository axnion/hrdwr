package hrdwr

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/**
 * Tests the RealRunner by running the echo command and validating the returned results
 */
func TestRealRunnerRun(t *testing.T) {
	runner = RealRunner{}
	results, err := runner.run("echo", "This is a test")

	assert.Nil(t, err)
	assert.Equal(t, "This is a test\n", string(results))
}