package lib

import (
	"strings"
	"strconv"
)

/**
 * Memory struct represents the RAM on the system. It has total amount of memory in the system and how much is currently
 * in use. Both numbers are in kilobytes
 */
type Memory struct {
	Total int
	Used int
}

/**
 * Returns a Memory object with the total amount of memory on the system and how much is currently in use. The data is
 * gathered from the /proc/meminfo file and then parsed into a Memory object.
 */
func GetMemory() (Memory, error) {
	content, err := runner.run("cat", "/proc/meminfo")
	if err != nil {
		var mem Memory
		return mem, err
	}

	return parseMeminfo(content)
}

/**
 * Takes the content of the /proc/meminfo file and parse it into a Memory object.
 */
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

			memory.Used = memory.Total - available
		}
	}

	return memory, nil
}