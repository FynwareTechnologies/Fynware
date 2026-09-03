package storage

type NVMe struct {
	Disk  *Disk
	Name  string
	Ready bool
}

func NewNVMe(size uint64) *NVMe {
	return &NVMe{
		Disk:  NewDisk(size),
		Name:  "Fynware NVMe Controller",
		Ready: true,
	}
}

func (n *NVMe) Capacity() uint64 {
	return n.Disk.Size()
}

func (n *NVMe) Read(address uint64, buffer []byte) error {
	return n.Disk.Read(address, buffer)
}

func (n *NVMe) Write(address uint64, data []byte) error {
	return n.Disk.Write(address, data)
}
