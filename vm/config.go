package vm

type Config struct {
	CPUArchitecture string
	CPUCores        int
	MemoryMB        uint64
	StorageGB       uint64
	EnablePCIe      bool
	EnableUSB       bool
	EnableNetwork   bool
	EnableACPI      bool
	EnableTPM       bool
}
