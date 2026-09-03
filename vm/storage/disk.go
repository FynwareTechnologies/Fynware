package storage

import "fmt"

type Disk struct {
	data []byte
}

func NewDisk(size uint64) *Disk {
	return &Disk{
		data: make([]byte, size),
	}
}

func (d *Disk) Size() uint64 {
	return uint64(len(d.data))
}

func (d *Disk) Read(address uint64, buffer []byte) error {
	if address+uint64(len(buffer)) > d.Size() {
		return fmt.Errorf("disk read out of bounds: 0x%X", address)
	}

	copy(buffer, d.data[address:address+uint64(len(buffer))])
	return nil
}

func (d *Disk) Write(address uint64, data []byte) error {
	if address+uint64(len(data)) > d.Size() {
		return fmt.Errorf("disk write out of bounds: 0x%X", address)
	}

	copy(d.data[address:address+uint64(len(data))], data)
	return nil
}
