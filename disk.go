package lib

import (
	"strings"
	"strconv"
)

/**
 * This struct represents a drive on the system. It has the name of the device, the total amount of storage
 * available on the device, and how much of that storage is currently being used.
 */
type Disk struct {
	Name string
	Total int
	Used int
}

/**
 * Returns a slice of Disk objects taken from the output of the "df" command and then parsed.
 */
func GetDisks() ([]Disk, error) {
	content, err := runner.run("df")

	if err != nil {
		return nil, err
	}

	return parseDf(content)
}

/**
 * Takes the output of the "df" command and parses it find hard drives in the current system and parses these into
 * Disk objects and return the slice containing them.
 */
func parseDf(content []byte) ([]Disk, error) {
	var disks []Disk
	str := string(content)
	lines:= strings.Split(str, "\n")

	for _, line := range lines {
		var drive []string
		columns := strings.Split(line, " ")

		if strings.Contains(columns[0], "/dev") {
			for _, column := range columns {
				if strings.Compare(column, "") != 0 {
					drive = append(drive, column)
				}
			}

			total, err := strconv.Atoi(drive[1])

			if err != nil {
				return nil, err
			}

			used, err := strconv.Atoi(drive[2])

			if err != nil {
				return nil, err
			}

			disks = append(disks, Disk{
				drive[0],
				total,
				used,
			})
		}
	}

	return disks, nil
}