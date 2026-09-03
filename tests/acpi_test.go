package tests

import (
	"testing"

	"github.com/FynwareTechnologies/Fynware/vm/acpi"
)

func TestACPI(t *testing.T) {
	system := acpi.New(true)

	if !system.Enabled {
		t.Fatal("ACPI should be enabled")
	}

	if system.Tables == nil {
		t.Fatal("ACPI tables should be initialized")
	}

	if system.Tables.RSDP == nil {
		t.Fatal("RSDP should be initialized")
	}

	if system.Tables.RSDT == nil {
		t.Fatal("RSDT should be initialized")
	}

	if system.Tables.FADT == nil {
		t.Fatal("FADT should be initialized")
	}

	if system.Tables.DSDT == nil {
		t.Fatal("DSDT should be initialized")
	}

	system.Stop()

	if system.Enabled {
		t.Fatal("ACPI should be disabled after Stop")
	}

	system.Start()

	if !system.Enabled {
		t.Fatal("ACPI should be enabled after Start")
	}
}
