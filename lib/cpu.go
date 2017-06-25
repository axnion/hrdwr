package lib

import (
	"time"
	"strings"
	"strconv"
)

/**
 * The CPU struct represents one CPU core. It has the name of the code, for example cpu0, and the usage. The usage is
 * measured so 0 = 0%, 1 = 100% 0.5 = 50%.
 */
type CPU struct {
	Name string
	Usage float64
}

/**
 * The procStat object is used to store information about one core as it is represented in the /proc/stat file. This
 * is to make management of the data easier when calculating the current usage of the CPU.
 */
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
 * Finds and returns all CPU cores and their current usage. This is code by reading the contence of the /proc/stat file
 * twice and with that data calculate the usage for each core.
 */
func GetCpus() ([]CPU, error) {
	var cpus []CPU

	content, err := runner.run("cat", "/proc/stat")
	if err != nil {return nil, err}
	stat1, err := parseProcStat(content)
	if err != nil {return nil, err}

	time.Sleep(500 * time.Millisecond)

	content, err = runner.run("cat", "/proc/stat")
	if err != nil {return nil, err}
	stat2, err := parseProcStat(content)
	if err != nil {return nil, err}

	for i := range stat1 {
		cpus = append(cpus, CPU{
			stat1[i].name,
			calcCpuUsage(stat1[i], stat2[i]),
		})
	}

	if err != nil {return nil, err}

	return cpus, nil
}


/**
 * Takes the content of a /proc/stat file and an array of CPU objects. It parses the file content and stores the data
 * in memory as procStat objects.
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
 * Calculates the utilization of a CPU by taking two reading of the /proc/stat file in the form of two procStat objects
 * and then comparing the data to calculate the total usage of the core.
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

