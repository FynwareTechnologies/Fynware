package firmware

import "fmt"

type Flash struct {
	data []byte
}

func NewFlash(size uint64) *Flash {
	return &Flash{
		data: make([]byte, size),
	}
}

func (f *Flash) Size() uint64 {
	return uint64(len(f.data))
}

func (f *Flash) Read(address uint64, buffer []byte) error {
	if address+uint64(len(buffer)) > f.Size() {
		return fmt.Errorf("flash read out of bounds: 0x%X", address)
	}

	copy(buffer, f.data[address:address+uint64(len(buffer))])
	return nil
}

func (f *Flash) Write(address uint64, data []byte) error {
	if address+uint64(len(data)) > f.Size() {
		return fmt.Errorf("flash write out of bounds: 0x%X", address)
	}

	copy(f.data[address:address+uint64(len(data))], data)
	return nil
}
