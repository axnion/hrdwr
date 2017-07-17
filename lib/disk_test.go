package lib

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

/*
 * - Test Runner -------------------------------------------------------------------------------------------------------
 */

/**
 * A struct which replaces the original Runner object used in the Disk test package.
 */
type DiskTestRunner struct {
	results []byte
	err error
}

/**
 * Mocked run method of DiskTestRunner struct. Takes the results field and err field and returns them as the result to
 * simulate running the "df" command
 */
func (runner *DiskTestRunner) run(cmd string, args ...string) ([]byte, error) {
	return runner.results, runner.err
}

/**
 * Replaces runner with DiskTestRunner so executions uses these mocks.
 */
func prepareDiskRunner(results []byte, err error) {
	runner = &DiskTestRunner{results, err}
}

/*
 * - Test Suite --------------------------------------------------------------------------------------------------------
 */

/**
 * Main test method. Runner returns a example results of a df command and then parses this string. The results is then
 * compared with the expected values to validate functionality.
 */
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

/**
 * Simulates what happens if command runner results in a error.
 */
func TestGetDisksError(t *testing.T) {
	errin := errors.New("test")
	prepareDiskRunner([]byte(""), errin)

	disk, err := GetDisks()

	assert.NotNil(t, err)
	assert.Nil(t, disk)
	assert.Equal(t, "test", err.Error())
}

/**
 * Running command return a results containing errors on the 1K-blocks column which inhibits parsing of the values.
 */
func TestGetDisksErrorParsingTotal(t *testing.T) {
	results := "Filesystem     1K-blocks     Used Available Use% Mounted on\n" +
		"dev              6144692        0   6144692   0% /dev\n" +
		"run              6149132     1968   6147164   1% /run\n" +
		"/dev/sdc2      2233uu90420 41044080 170929048  20% /\n" +
		"tmpfs            6149132   100004   6049128   2% /dev/shm\n" +
		"tmpfs            6149132        0   6149132   0% /sys/fs/cgroup\n" +
		"/dev/sdc1         248775    48077    183546  21% /boot\n" +
		"tmpfs            6149132     3068   6146064   1% /tmp\n" +
		"tmpfs            1229824       40   1229784   1% /run/user/1000\n"

	prepareDiskRunner([]byte(results), nil)

	disk, err := GetDisks()

	assert.NotNil(t, err)
	assert.Nil(t, disk)
}

/**
 * Running command return a results containing errors on the Used column which inhibits parsing of the values.
 */
func TestGetDisksErrorParsingUsed(t *testing.T) {
	results := "Filesystem     1K-blocks     Used Available Use% Mounted on\n" +
		"dev              6144692        0   6144692   0% /dev\n" +
		"run              6149132     1968   6147164   1% /run\n" +
		"/dev/sdc2      223390420 4104ii4080 170929048  20% /\n" +
		"tmpfs            6149132   100004   6049128   2% /dev/shm\n" +
		"tmpfs            6149132        0   6149132   0% /sys/fs/cgroup\n" +
		"/dev/sdc1         248775    48077    183546  21% /boot\n" +
		"tmpfs            6149132     3068   6146064   1% /tmp\n" +
		"tmpfs            1229824       40   1229784   1% /run/user/1000\n"

	prepareDiskRunner([]byte(results), nil)

	disk, err := GetDisks()

	assert.NotNil(t, err)
	assert.Nil(t, disk)
}
