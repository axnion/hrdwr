package units

import (
	"github.com/axnion/hrdwr/util"
	"strings"
	"fmt"
	"strconv"
)

type CpuMon struct{}

type CPU struct {
	id string
	usage float32
}

var runner util.Runner

func (cpu CpuMon) GetCpus() ([]CPU, error) {
	stat, _ := run(runner, "cat", "/proc/stat")

	parseProcStat(stat)

	return nil, nil
}

func parseProcStat(content []byte, cpus []CPU) {
	str := string(content)
	lines := strings.Split(str, "\n")

	for i, line := range lines {
		columns := strings.Split(line, " ")

		if strings.Compare(columns[0], "cpu" + strconv.Itoa(i - 1)) == 0 {
			fmt.Println(line)
		}
	}
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
