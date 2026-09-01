package vm

import (
	"fmt"

	"github.com/FynwareTechnologies/Fynware/vm/cpu"
	"github.com/FynwareTechnologies/Fynware/vm/memory"
)

type VM struct {
	Config Config
	CPU    *cpu.CPU
	Memory *memory.Memory
}

func New(config Config) *VM {
	return &VM{
		Config: config,
		CPU:    cpu.New(config.CPUArchitecture, config.CPUCores),
		Memory: memory.New(config.MemoryMB * 1024 * 1024),
	}
}

func (v *VM) Start() {
	v.CPU.Start()

	fmt.Printf("CPU:       %s (%d cores)\n", v.Config.CPUArchitecture, v.Config.CPUCores)
	fmt.Printf("Memory:    %d MB\n", v.Config.MemoryMB)
	fmt.Printf("Storage:   %d GB\n", v.Config.StorageGB)
	fmt.Printf("PCIe:      %t\n", v.Config.EnablePCIe)
	fmt.Printf("USB:       %t\n", v.Config.EnableUSB)
	fmt.Printf("Network:   %t\n", v.Config.EnableNetwork)
	fmt.Printf("ACPI:      %t\n", v.Config.EnableACPI)
	fmt.Printf("TPM:       %t\n", v.Config.EnableTPM)
	fmt.Println()
	fmt.Println("Virtual machine started.")
}
