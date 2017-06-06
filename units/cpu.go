package units

import (
	"github.com/axnion/hrdwr/util"
	"strings"
	"strconv"
)

type CpuMon struct{}

type CPU struct {
	Name string

	Usage float64
	idle int
	total int

}

var runner util.Runner
var cpus = []CPU{}

func (cpu CpuMon) GetCpus(cpus []CPU) ([]CPU, error) {
	stat, _ := run(runner, "cat", "/proc/stat")

	return parseProcStat(stat, cpus)
}

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

func run(runner util.Runner, command string, arg string) ([]byte, error) {
	return runner.Run(command, arg)
}

//func parseCpuinfo(content []byte) ([]Core) {

//}

//func getCpuinfo(runner util.Runner) ([]byte, error) {
//	return runner.Run("cat", "/proc/cpuinfo")
//}

func (cpu CpuMon) SetRunner(newRunner util.Runner) {
	runner = newRunner
}
