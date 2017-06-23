package lib

import (
	"strings"
	"strconv"
)

//var runner Runner = RealRunner{}

type Disk struct {
	Name string
	Total int
	Used int
}

func GetDisks() ([]Disk, error) {
	content, err := runner.run("df")

	if err != nil {
		return nil, err
	}

	return parseDf(content)
}

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