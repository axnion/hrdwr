package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
 * - Test Runner -----------------------------------------------------------------------------------
 */

type DiskTestRunner struct {
	results []byte
	err error
}

func (runner *DiskTestRunner) run(cmd string, args ...string) ([]byte, error) {
	return runner.results, runner.err
}

/*
 * - Test Suite ------------------------------------------------------------------------------------
 */

func prepareDiskRunner(results []byte, err error) {
	runner = &DiskTestRunner{results, err}
}

func TestGetDisks(t *testing.T) {
	results := "Filesystem     1K-blocks     Used Available Use% Mounted on\n" +
		"dev              6144692        0   6144692   0% /dev\n" +
		"run              6149132     1968   6147164   1% /run\n" +
		"/dev/sdc2      223390420 41044080 170929048  20% /\n" +
		"tmpfs            6149132   100004   6049128   2% /dev/shm\n" +
		"tmpfs            6149132        0   6149132   0% /sys/fs/cgroup\n" +
		"/dev/sdc1         248775    48077    183546  21% /boot\n" +
		"tmpfs            6149132     3068   6146064   1% /tmp\n" +
		"tmpfs            1229824       40   1229784   1% /run/user/1000\n"

	prepareDiskRunner([]byte(results), nil)

	disk, err := GetDisks()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(disk))
	assert.Equal(t, disk[0].Name, "/dev/sdc2")
	assert.Equal(t, disk[1].Name, "/dev/sdc1")
	assert.Equal(t, disk[0].Total, 223390420)
	assert.Equal(t, disk[1].Total, 248775)
	assert.Equal(t, disk[0].Used, 41044080)
	assert.Equal(t, disk[1].Used, 48077)
}