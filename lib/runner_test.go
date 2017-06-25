package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRealRunnerRun(t *testing.T) {
	runner = RealRunner{}
	results, err := runner.run("echo", "This is a test")

	assert.Nil(t, err)
	assert.Equal(t, "This is a test\n", string(results))
}