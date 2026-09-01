package tests

import (
	"testing"

	"github.com/FynwareTechnologies/Fynware/vm"
)

func TestVMCreation(t *testing.T) {
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

	if machine == nil {
		t.Fatal("expected VM to be created")
	}

	if machine.Config.CPUArchitecture != "x86_64" {
		t.Fatalf("expected x86_64 architecture, got %s", machine.Config.CPUArchitecture)
	}

	if machine.Config.CPUCores != 4 {
		t.Fatalf("expected 4 CPU cores, got %d", machine.Config.CPUCores)
	}

	if machine.Config.MemoryMB != 8192 {
		t.Fatalf("expected 8192 MB memory, got %d", machine.Config.MemoryMB)
	}

	if machine.Config.StorageGB != 64 {
		t.Fatalf("expected 64 GB storage, got %d", machine.Config.StorageGB)
	}
}
