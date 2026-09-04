package tests

import (
	"testing"

	"github.com/FynwareTechnologies/Fynware/vm/display"
)

func TestDisplay(t *testing.T) {
	screen := display.New(1920, 1080)

	if screen.Width != 1920 {
		t.Fatalf("expected width 1920, got %d", screen.Width)
	}

	if screen.Height != 1080 {
		t.Fatalf("expected height 1080, got %d", screen.Height)
	}

	if !screen.Enabled {
		t.Fatal("display should be enabled")
	}

	screen.Stop()

	if screen.Enabled {
		t.Fatal("display should be disabled after Stop")
	}

	screen.Start()

	if !screen.Enabled {
		t.Fatal("display should be enabled after Start")
	}
}

func TestGPU(t *testing.T) {
	gpu := display.NewGPU("Fynware Virtual GPU")

	if gpu.Name != "Fynware Virtual GPU" {
		t.Fatalf("unexpected GPU name: %s", gpu.Name)
	}

	if gpu.Enabled {
		t.Fatal("GPU should be disabled initially")
	}

	gpu.Start()

	if !gpu.Enabled {
		t.Fatal("GPU should be enabled after Start")
	}

	gpu.Stop()

	if gpu.Enabled {
		t.Fatal("GPU should be disabled after Stop")
	}
}
