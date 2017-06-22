package parse

import (
	"strings"
	"strconv"
)

/**
 * A row from /proc/stat represented in memory.
 */
type procStat struct {
	Name string
	user int
	nice int
	system int
	idle int
	iowait int
	irq int
	softirq int
	steal int
}

/**
 * Takes the content of the /proc/stat file and an array of CPU objects. It parses the file content
 * and calculates the cpu usage. The data is then stored in the CPU array.
 */
func (RealParser) ProcStat(content []byte) ([]procStat, error) {
	var stat []procStat
	str := string(content)
	lines := strings.Split(str, "\n")

	for index, line := range lines {
		i := index -1
		columns := strings.Split(line, " ")

		if strings.Compare(columns[0], "cpu" + strconv.Itoa(i)) == 0 {
			user, err := strconv.Atoi(columns[1])
			if err != nil {return nil, err}

			nice, err := strconv.Atoi(columns[2])
			if err != nil {return nil, err}

			system, err := strconv.Atoi(columns[3])
			if err != nil {return nil, err}

			idle, err := strconv.Atoi(columns[4])
			if err != nil {return nil, err}

			iowait, err := strconv.Atoi(columns[5])
			if err != nil {return nil, err}

			irq, err := strconv.Atoi(columns[6])
			if err != nil {return nil, err}

			softirq, err := strconv.Atoi(columns[7])
			if err != nil {return nil, err}

			steal, err := strconv.Atoi(columns[8])
			if err != nil {return nil, err}

			stat = append(stat, procStat{
				Name: columns[0],
				user: user,
				nice: nice,
				system: system,
				idle: idle,
				iowait: iowait,
				irq: irq,
				softirq: softirq,
				steal: steal,
			})
		}
	}

	return stat, nil
}

/**
 * Calculates the CPU utilization based on two readings of /proc/stat
 */
func (RealParser) CalcCpuUsage(prev procStat, cur procStat) float64 {
	prevIdle := prev.idle + prev.iowait
	idle := cur.idle + cur.iowait

	prevNonIdle := prev.user + prev.nice + prev.system + prev.irq + prev.softirq + prev.steal
	nonIdle := cur.user + cur.nice + cur.system + cur.irq + cur.softirq + cur.steal

	prevTotal := prevIdle + prevNonIdle
	total := idle + nonIdle

	totalDiff := total - prevTotal
	idleDiff := idle - prevIdle

	return float64(totalDiff - idleDiff) / float64(totalDiff)
}

