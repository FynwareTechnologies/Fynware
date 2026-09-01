package cpu

type Instruction struct {
	Opcode []byte
	Length int
}

func (c *CPU) Execute(instruction Instruction) {
	if !c.Running {
		return
	}

	c.Registers.RIP += uint64(instruction.Length)
}
