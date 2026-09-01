package cpu

type CPU struct {
	Architecture string
	Cores        int
	Registers    Registers
	Running      bool
}

func New(architecture string, cores int) *CPU {
	return &CPU{
		Architecture: architecture,
		Cores:        cores,
	}
}

func (c *CPU) Reset() {
	c.Registers = Registers{}
	c.Running = false
}

func (c *CPU) Start() {
	c.Running = true
}

func (c *CPU) Stop() {
	c.Running = false
}
