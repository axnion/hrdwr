package units

import (
	"github.com/axnion/hrdwr/util"
	"strings"
	"strconv"
)

/**
 * CpuMon (CPU Monitoring) is the object which holds the methods for fetching data on the CPU
 */
type CpuMon struct{
	cpus []CPU
 	runner util.Runner
}

/**
 * Final representation of a CPU
 */
type CPU struct {
	Name string

	Usage float64
	idle int
	total int
}

func NewCpuMon(runner util.Runner) CpuMon {
	mon := CpuMon{
		cpus: []CPU{},
		runner: runner,
	}

	return mon
}

func (mon CpuMon) SetRunner(newRunner util.Runner) {
	mon.runner = newRunner
}
/**
 * Returns an array of CPU objects which represents the CPUs of the system.
 */
func (mon CpuMon) GetCpus() ([]CPU, error) {
	stat, _ := run(mon.runner, "cat", "/proc/stat")

	cpus, err := parseProcStat(stat, mon.cpus)

	if err != nil {
		return nil, err
	}

	return cpus, nil
}

/**
 * Takes the content of the /proc/stat file and an array of CPU objects. It parses the file content
 * and calculates the cpu usage. The data is then stored in the CPU array.
 */
func parseProcStat(content []byte, cpus []CPU) ([]CPU, error) {
	str := string(content)
	lines := strings.Split(str, "\n")

	for index, line := range lines {
		i := index -1
		columns := strings.Split(line, " ")

		if strings.Compare(columns[0], "cpu" + strconv.Itoa(i)) == 0 {
			if len(cpus) == i {
				cpus = append(cpus, CPU{
					Name: columns[0],
					idle: 0,
					total: 0,
				})
			}

			user, err := strconv.Atoi(columns[1])

			if err != nil {
				return nil, err
			}

			nice, err := strconv.Atoi(columns[2])

			if err != nil {
				return nil, err
			}

			system, err := strconv.Atoi(columns[3])

			if err != nil {
				return nil, err
			}

			idle, err := strconv.Atoi(columns[4])

			if err != nil {
				return nil, err
			}

			iowait, err := strconv.Atoi(columns[5])

			if err != nil {
				return nil, err
			}

			irq, err := strconv.Atoi(columns[6])

			if err != nil {
				return nil, err
			}

			softirq, err := strconv.Atoi(columns[7])

			if err != nil {
				return nil, err
			}

			steal, err := strconv.Atoi(columns[8])

			if err != nil {
				return nil, err
			}

			newIdle := idle + iowait
			newTotal := newIdle + user + nice + system + irq + softirq + steal

			totalDiff := newTotal - cpus[i].total
			idleDiff := newIdle - cpus[i].idle

			usage := float64(totalDiff - idleDiff) / float64(totalDiff)

			cpus[i].idle = newIdle
			cpus[i].total = newTotal
			cpus[i].Usage = usage
		}
	}

	return cpus, nil
}

/**
 * Takes a runner, a command string, and an arguments string. It runs the command using the runner
 * and the argument.
 */
func run(runner util.Runner, command string, arg string) ([]byte, error) {
	return runner.Run(command, arg)
}
