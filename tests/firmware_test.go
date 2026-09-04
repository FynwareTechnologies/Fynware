package tests

import (
	"bytes"
	"testing"

	"github.com/FynwareTechnologies/Fynware/vm/firmware"
)

func TestFirmware(t *testing.T) {
	image := []byte{0xF1, 0x4E, 0x57, 0x01}
	fw := firmware.New("Fynware Firmware", "0.1.0", image)

	if fw.Name != "Fynware Firmware" {
		t.Fatalf("unexpected firmware name: %s", fw.Name)
	}

	if fw.Version != "0.1.0" {
		t.Fatalf("unexpected firmware version: %s", fw.Version)
	}

	if !bytes.Equal(fw.Image, image) {
		t.Fatal("firmware image does not match")
	}

	if fw.Loaded {
		t.Fatal("firmware should not be loaded initially")
	}

	fw.Load()

	if !fw.Loaded {
		t.Fatal("firmware should be loaded after Load")
	}

	fw.Unload()

	if fw.Loaded {
		t.Fatal("firmware should be unloaded after Unload")
	}
}

func TestFlash(t *testing.T) {
	flash := firmware.NewFlash(1024)

	if flash.Size() != 1024 {
		t.Fatalf("expected flash size 1024, got %d", flash.Size())
	}

	data := []byte{0xF1, 0x4E, 0x57, 0x01}

	if err := flash.Write(100, data); err != nil {
		t.Fatalf("flash write failed: %v", err)
	}

	buffer := make([]byte, len(data))

	if err := flash.Read(100, buffer); err != nil {
		t.Fatalf("flash read failed: %v", err)
	}

	if !bytes.Equal(buffer, data) {
		t.Fatal("flash data does not match")
	}

	if err := flash.Write(1021, data); err == nil {
		t.Fatal("expected out-of-bounds flash write to fail")
	}

	if err := flash.Read(1021, buffer); err == nil {
		t.Fatal("expected out-of-bounds flash read to fail")
	}
}
