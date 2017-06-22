package libs

import (
	"time"
	"github.com/axnion/hrdwr/libs/parse"
)

/**
 * CpuMon (CPU Monitoring) is the object which holds the methods for fetching data on the CPU
 */
type CpuMon struct{
 	runner Runner
	parser parse.Parser
}


/**
 * Final representation of a CPU
 */
type CPU struct {
	Name string
	Usage float64
}

/**
 * Constructor function
 */
func NewCpuMon(runner Runner, parser parse.Parser) CpuMon {
	return CpuMon{
		runner: runner,
		parser: parser,
	}
}

/**
 * Returns an array of CPU objects which represents the CPUs of the system.
 */
func (mon CpuMon) GetCpus() ([]CPU, error) {
	var cpus []CPU

	content, _ := run(mon.runner, "cat", "/proc/stat")
	stat1, err := mon.parser.ProcStat(content)

	time.Sleep(500 * time.Millisecond)

	content, _ = run(mon.runner, "cat", "/proc/stat")
	stat2, err := mon.parser.ProcStat(content)


	for i := range stat1 {
		cpus = append(cpus, CPU{
			stat1[i].Name,
			mon.parser.CalcCpuUsage(stat1[i], stat2[i]),
		})
	}

	if err != nil {
		return nil, err
	}

	return cpus, nil
}
