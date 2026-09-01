package cpu

type Registers struct {
	RAX uint64
	RBX uint64
	RCX uint64
	RDX uint64

	RSI uint64
	RDI uint64
	RBP uint64
	RSP uint64

	R8  uint64
	R9  uint64
	R10 uint64
	R11 uint64
	R12 uint64
	R13 uint64
	R14 uint64
	R15 uint64

	RIP    uint64
	RFLAGS uint64
}
