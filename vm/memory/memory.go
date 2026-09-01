package memory

import "fmt"

type Memory struct {
	data []byte
}

func New(size uint64) *Memory {
	return &Memory{
		data: make([]byte, size),
	}
}

func (m *Memory) Size() uint64 {
	return uint64(len(m.data))
}

func (m *Memory) Read8(address Address) (byte, error) {
	if uint64(address) >= uint64(len(m.data)) {
		return 0, fmt.Errorf("memory read out of bounds: 0x%X", address)
	}

	return m.data[address], nil
}

func (m *Memory) Write8(address Address, value byte) error {
	if uint64(address) >= uint64(len(m.data)) {
		return fmt.Errorf("memory write out of bounds: 0x%X", address)
	}

	m.data[address] = value
	return nil
}
