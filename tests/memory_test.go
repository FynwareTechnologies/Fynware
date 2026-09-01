package tests

import (
	"testing"

	"github.com/FynwareTechnologies/Fynware/vm/memory"
)

func TestMemoryCreation(t *testing.T) {
	m := memory.New(1024)

	if m.Size() != 1024 {
		t.Fatalf("expected memory size 1024, got %d", m.Size())
	}
}

func TestMemoryReadWrite(t *testing.T) {
	m := memory.New(1024)

	value := byte(0x42)

	if err := m.Write8(0x100, value); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	got, err := m.Read8(0x100)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}

	if got != value {
		t.Fatalf("expected 0x%02X, got 0x%02X", value, got)
	}
}

func TestMemoryOutOfBounds(t *testing.T) {
	m := memory.New(1024)

	if err := m.Write8(1024, 0x42); err == nil {
		t.Fatal("expected out-of-bounds write to fail")
	}

	if _, err := m.Read8(1024); err == nil {
		t.Fatal("expected out-of-bounds read to fail")
	}
}
