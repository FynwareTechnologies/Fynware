package platform

type Hardware struct {
	CPUArchitecture string
	CPUCores        int
	MemoryMB        uint64
	StorageGB       uint64

	PCIeEnabled    bool
	USBEnabled     bool
	NetworkEnabled bool
	ACPIEnabled    bool
	TPMEnabled     bool
}

func NewHardware(
	cpuArchitecture string,
	cpuCores int,
	memoryMB uint64,
	storageGB uint64,
	pcieEnabled bool,
	usbEnabled bool,
	networkEnabled bool,
	acpiEnabled bool,
	tpmEnabled bool,
) *Hardware {
	return &Hardware{
		CPUArchitecture: cpuArchitecture,
		CPUCores:        cpuCores,
		MemoryMB:        memoryMB,
		StorageGB:       storageGB,
		PCIeEnabled:     pcieEnabled,
		USBEnabled:      usbEnabled,
		NetworkEnabled:  networkEnabled,
		ACPIEnabled:     acpiEnabled,
		TPMEnabled:      tpmEnabled,
	}
}
