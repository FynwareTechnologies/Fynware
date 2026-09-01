package main

import (
	"fmt"

	"github.com/FynwareTechnologies/Fynware/vm"
)

func main() {
	config := vm.Config{
		CPUArchitecture: "x86_64",
		CPUCores:        4,
		MemoryMB:        8192,
		StorageGB:       64,
		EnablePCIe:      true,
		EnableUSB:       true,
		EnableNetwork:   true,
		EnableACPI:      true,
		EnableTPM:       true,
	}

	machine := vm.New(config)

	fmt.Println("Fynware VM")
	fmt.Println("==========")
	machine.Start()
}
