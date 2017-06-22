package lib

import (
	"strings"
	"strconv"
)

type MemMon struct {
	runner Runner
}

type Memory struct {
	Total int
	Available int	// TODO: Change to used memory
}

func NewMemMon(runner Runner) MemMon {
	return MemMon{
		runner: runner,
	}
}

func (mon MemMon) GetMemory() (Memory, error) {
	content, _ := run(mon.runner, "cat", "/proc/meminfo")
	return parseMeminfo(content)
}

func parseMeminfo(content []byte) (Memory, error) {
	var memory Memory
	str := string(content)
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		column := strings.Split(line, " ")

		if strings.Contains(line, "MemTotal") {
			total, err := strconv.Atoi(column[len(column) - 2])

			if err != nil {
				return memory, err
			}

			memory.Total= total
		}

		if strings.Contains(line, "MemAvailable") {
			available, err := strconv.Atoi(column[len(column) - 2])

			if err != nil {
				return memory, err
			}

			memory.Available = available
		}
	}

	return memory, nil
}