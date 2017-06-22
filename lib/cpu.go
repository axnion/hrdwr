package lib

import (
	"time"
	"strings"
	"strconv"
)

var runner Runner = RealRunner{}

/**
 * Final representation of a CPU
 */
type CPU struct {
	Name string
	Usage float64
}

type procStat struct {
	name string
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
 * Returns an array of CPU objects which represents the CPUs of the system.
 */
func GetCpus() ([]CPU, error) {
	var cpus []CPU

	content, _ := run(runner, "cat", "/proc/stat")
	stat1, err := parseProcStat(content)

	time.Sleep(500 * time.Millisecond)

	content, _ = run(runner, "cat", "/proc/stat")
	stat2, err := parseProcStat(content)


	for i := range stat1 {
		cpus = append(cpus, CPU{
			stat1[i].name,
			calcCpuUsage(stat1[i], stat2[i]),
		})
	}

	if err != nil {
		return nil, err
	}

	return cpus, nil
}


/**
 * Takes the content of the /proc/stat file and an array of CPU objects. It parses the file content
 * and calculates the cpu usage. The data is then stored in the CPU array.
 */
func parseProcStat(content []byte) ([]procStat, error) {
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
				name: columns[0],
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
func calcCpuUsage(prev procStat, cur procStat) float64 {
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

