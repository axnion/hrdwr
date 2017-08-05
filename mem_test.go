package hrdwr

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"
)

/*
 * - Test Runner -------------------------------------------------------------------------------------------------------
 */

/**
 * MemTestRunner struct to replace the runner to mock the execution of commands.
 */
type MemTestRunner struct {
	results []byte
	err error
}

/**
 * Mocked run method for MemTestRunner. Returns the fields results and error instead of executing the command.
 */
func (runner *MemTestRunner) run(cmd string, args ...string) ([]byte, error) {
	return runner.results, runner.err
}

/**
 * Replaces the runner object with a MemTestRunner pointer.
 */
func prepareMemRunner(results []byte, err error) {
	runner = &MemTestRunner{results, err}
}

/*
 * - Test Suite --------------------------------------------------------------------------------------------------------
 */

/**
 * The main test method for memory. An example result from /proc/meminfo is returned when run command is executed
 * thanks to MemTestRunner. Results are then compared with the expected values.
 */
func TestGetMemory(t *testing.T) {
	results := "MemTotal:       12298264 kB\nMemFree:         4887148 kB\nMemAvailable:    6775688 kB\n" +
		"Buffers:           93396 kB\nCached:          2334216 kB\nSwapCached:            0 kB\n" +
		"Active:          5292924 kB\nInactive:        1560276 kB\nActive(anon):    4429052 kB\n" +
		"Inactive(anon):   368268 kB\nActive(file):     863872 kB\nInactive(file):  1192008 kB\n" +
		"Unevictable:         172 kB\nMlocked:             172 kB\nSwapTotal:       6149716 kB\n" +
		"SwapFree:        6149716 kB\nDirty:             11852 kB\nWriteback:             0 kB\n" +
		"AnonPages:       4277552 kB\nMapped:          1239376 kB\nShmem:            371736 kB\n" +
		"Slab:             233028 kB\nSReclaimable:     140712 kB\nSUnreclaim:        92316 kB\n" +
		"KernelStack:       17024 kB\nPageTables:        67024 kB\nNFS_Unstable:          0 kB\n" +
		"Bounce:                0 kB\nWritebackTmp:          0 kB\nCommitLimit:    12298848 kB\n" +
		"Committed_AS:   12639456 kB\nVmallocTotal:   34359738367 kB\nVmallocUsed:           0 kB\n" +
		"VmallocChunk:          0 kB\nHardwareCorrupted:     0 kB\nAnonHugePages:   1697792 kB\n" +
		"ShmemHugePages:        0 kB\nShmemPmdMapped:        0 kB\nHugePages_Total:       0\n" +
		"HugePages_Free:        0\nHugePages_Rsvd:        0\nHugePages_Surp:        0\n" +
		"Hugepagesize:       2048 kB\nDirectMap4k:      370176 kB\nDirectMap2M:    12204032 kB\n"

	prepareMemRunner([]byte(results), nil)

	memory, err := GetMemory()

	assert.Nil(t, err)
	assert.NotNil(t, memory)
	assert.Equal(t, 12298264, memory.Total)
	assert.Equal(t, 5522576, memory.Used)
}

/**
 * Simulates if the command would result in an error.
 */
func TestGetMemoryRunError(t *testing.T) {
	errin := errors.New("test")
	prepareMemRunner([]byte(""), errin)

	memory, err := GetMemory()

	assert.NotNil(t, err)
	assert.Equal(t, "test", err.Error())
	assert.NotNil(t, memory)
	assert.Equal(t, 0, memory.Total)
	assert.Equal(t, 0, memory.Used)
}

/**
 * Tests how parsing is handled when MemTotal columns has an error
 */
func TestGetMemoryTotalParseError(t *testing.T) {
	results := "MemTotal:       1229i8264 kB\nMemFree:         4887148 kB\nMemAvailable:    6775688 kB\n" +
		"Buffers:           93396 kB\nCached:          2334216 kB\nSwapCached:            0 kB\n" +
		"Active:          5292924 kB\nInactive:        1560276 kB\nActive(anon):    4429052 kB\n" +
		"Inactive(anon):   368268 kB\nActive(file):     863872 kB\nInactive(file):  1192008 kB\n" +
		"Unevictable:         172 kB\nMlocked:             172 kB\nSwapTotal:       6149716 kB\n" +
		"SwapFree:        6149716 kB\nDirty:             11852 kB\nWriteback:             0 kB\n" +
		"AnonPages:       4277552 kB\nMapped:          1239376 kB\nShmem:            371736 kB\n" +
		"Slab:             233028 kB\nSReclaimable:     140712 kB\nSUnreclaim:        92316 kB\n" +
		"KernelStack:       17024 kB\nPageTables:        67024 kB\nNFS_Unstable:          0 kB\n" +
		"Bounce:                0 kB\nWritebackTmp:          0 kB\nCommitLimit:    12298848 kB\n" +
		"Committed_AS:   12639456 kB\nVmallocTotal:   34359738367 kB\nVmallocUsed:           0 kB\n" +
		"VmallocChunk:          0 kB\nHardwareCorrupted:     0 kB\nAnonHugePages:   1697792 kB\n" +
		"ShmemHugePages:        0 kB\nShmemPmdMapped:        0 kB\nHugePages_Total:       0\n" +
		"HugePages_Free:        0\nHugePages_Rsvd:        0\nHugePages_Surp:        0\n" +
		"Hugepagesize:       2048 kB\nDirectMap4k:      370176 kB\nDirectMap2M:    12204032 kB\n"

	prepareMemRunner([]byte(results), nil)

	memory, err := GetMemory()

	assert.NotNil(t, err)
	assert.NotNil(t, memory)
	assert.Equal(t, 0, memory.Total)
	assert.Equal(t, 0, memory.Used)
}

/**
 * Tests how parsing is handled when MemAvailable columns has an error
 */
func TestGetMemoryAvailableParseError(t *testing.T) {
	results := "MemTotal:       12298264 kB\nMemFree:         4887148 kB\nMemAvailable:    6775a688 kB\n" +
		"Buffers:           93396 kB\nCached:          2334216 kB\nSwapCached:            0 kB\n" +
		"Active:          5292924 kB\nInactive:        1560276 kB\nActive(anon):    4429052 kB\n" +
		"Inactive(anon):   368268 kB\nActive(file):     863872 kB\nInactive(file):  1192008 kB\n" +
		"Unevictable:         172 kB\nMlocked:             172 kB\nSwapTotal:       6149716 kB\n" +
		"SwapFree:        6149716 kB\nDirty:             11852 kB\nWriteback:             0 kB\n" +
		"AnonPages:       4277552 kB\nMapped:          1239376 kB\nShmem:            371736 kB\n" +
		"Slab:             233028 kB\nSReclaimable:     140712 kB\nSUnreclaim:        92316 kB\n" +
		"KernelStack:       17024 kB\nPageTables:        67024 kB\nNFS_Unstable:          0 kB\n" +
		"Bounce:                0 kB\nWritebackTmp:          0 kB\nCommitLimit:    12298848 kB\n" +
		"Committed_AS:   12639456 kB\nVmallocTotal:   34359738367 kB\nVmallocUsed:           0 kB\n" +
		"VmallocChunk:          0 kB\nHardwareCorrupted:     0 kB\nAnonHugePages:   1697792 kB\n" +
		"ShmemHugePages:        0 kB\nShmemPmdMapped:        0 kB\nHugePages_Total:       0\n" +
		"HugePages_Free:        0\nHugePages_Rsvd:        0\nHugePages_Surp:        0\n" +
		"Hugepagesize:       2048 kB\nDirectMap4k:      370176 kB\nDirectMap2M:    12204032 kB\n"

	prepareMemRunner([]byte(results), nil)

	memory, err := GetMemory()

	assert.NotNil(t, err)
	assert.NotNil(t, memory)
	assert.Equal(t, 12298264, memory.Total)
	assert.Equal(t, 0, memory.Used)
}
