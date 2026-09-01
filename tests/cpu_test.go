package tests

import (
	"testing"

	"github.com/FynwareTechnologies/Fynware/vm/cpu"
)

func TestCPUCreation(t *testing.T) {
	c := cpu.New("x86_64", 4)

	if c.Architecture != "x86_64" {
		t.Fatalf("expected x86_64, got %s", c.Architecture)
	}

	if c.Cores != 4 {
		t.Fatalf("expected 4 cores, got %d", c.Cores)
	}

	if c.Running {
		t.Fatal("new CPU should not be running")
	}
}

func TestCPUStartStop(t *testing.T) {
	c := cpu.New("x86_64", 4)

	c.Start()

	if !c.Running {
		t.Fatal("CPU should be running after Start")
	}

	c.Stop()

	if c.Running {
		t.Fatal("CPU should not be running after Stop")
	}
}

func TestCPUReset(t *testing.T) {
	c := cpu.New("x86_64", 4)

	c.Registers.RAX = 1234
	c.Registers.RIP = 0x1000
	c.Start()

	c.Reset()

	if c.Registers.RAX != 0 {
		t.Fatal("RAX should be zero after reset")
	}

	if c.Registers.RIP != 0 {
		t.Fatal("RIP should be zero after reset")
	}

	if c.Running {
		t.Fatal("CPU should not be running after reset")
	}
}
