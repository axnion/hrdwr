package libs

type DiskMon struct {
	runner Runner
}

type Disk struct {
	Name string
	Total int
	Used int
}

func NewDiskMon(runner Runner) DiskMon{
	return DiskMon{
		runner: runner,
	}
}

func (mon DiskMon) GetDisks() ([]Disk, error) {


	return nil, nil
}