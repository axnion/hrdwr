package parse

type Parser interface {
	ProcStat([]byte) ([]procStat, error)
	CalcCpuUsage(procStat, procStat) float64
}

type RealParser struct {}

type TestParser struct {
	parseProcStatExecutions int
}

func (parser TestParser) ParseProcStat(content []byte) ([]procStat, error) {
	parser.parseProcStatExecutions++
	var stat []procStat

		stat = []procStat{
			{
				"cpu0",
				136139,
				222,
				23664,
				1477316,
				2617,
				0,
				5331,
				0,

			},
		}

	return stat, nil
}

func (TestParser) CalcCpuUsage(stat1 procStat, stat2 procStat) float64 {
	return float64(0.7)
}

