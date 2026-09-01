# Fynware

**Virtual hardware and development infrastructure for building, testing, and validating firmware.**

Fynware is a firmware technology project focused on providing the infrastructure needed to develop, test, and validate firmware before it reaches real hardware.

The project is being built from the ground up, starting with **Fynware VM** — a virtual hardware environment designed specifically for firmware development and testing.

## Fynware VM

Fynware VM provides a programmable virtual hardware platform that firmware can eventually run against.

The virtual platform is designed to model hardware commonly found in modern computers, including:

* x86_64 CPUs
* System memory
* Storage
* PCIe
* USB
* Networking
* Display and GPU hardware
* ACPI
* TPM
* SPI flash
* Platform and motherboard hardware

The goal is to make firmware development possible without requiring physical hardware for every test.

## Project Structure

```text
Fynware/
├── cmd/
│   └── fynware-vm/       # Fynware VM executable
├── tests/                 # VM and subsystem tests
├── tools/                 # Development and inspection tools
└── vm/                    # Virtual hardware implementation
    ├── acpi/
    ├── cpu/
    ├── debug/
    ├── devices/
    ├── display/
    ├── firmware/
    ├── memory/
    ├── network/
    ├── platform/
    └── storage/
```

## Current Status

Fynware VM is currently under active development.

The current foundation includes:

* VM configuration
* x86_64 CPU model
* CPU registers
* CPU execution framework
* Virtual memory
* Memory addressing
* VM startup
* Automated tests

The virtual CPU is currently a development framework rather than a complete x86-64 emulator. Instruction execution will be expanded as the VM develops.

## Development

Fynware is written primarily in **Go**.

Run the VM:

```bash
go run ./cmd/fynware-vm
```

Run the tests:

```bash
go test ./tests
```

As additional subsystems are implemented, the test suite will expand to cover the complete virtual platform.

## Roadmap

### Fynware VM

* [x] VM foundation
* [x] CPU model
* [x] Memory model
* [ ] Complete CPU instruction execution
* [ ] Storage devices
* [ ] PCIe devices
* [ ] USB devices
* [ ] Network devices
* [ ] Display/GPU model
* [ ] ACPI tables
* [ ] TPM model
* [ ] SPI flash
* [ ] Platform/motherboard model
* [ ] Firmware loading
* [ ] Firmware execution
* [ ] VM debugging and tracing

### Fynware Firmware

Once the virtual hardware platform is sufficiently complete, Fynware Firmware will be developed and tested against Fynware VM.

The long-term goal is to allow firmware to be developed and validated in a controlled virtual environment before being deployed to physical hardware.

## Philosophy

> **Build the hardware environment first. Then build the firmware for it.**

Fynware is intended to make firmware development more accessible, testable, and repeatable.

Instead of immediately depending on physical computers, developers can use a virtual hardware platform to reproduce hardware configurations, test firmware behaviour, and investigate failures.

## License

Fynware is proprietary software.

See [`LICENSE`](LICENSE) for the terms governing use of this software and its source code.

---

**Fynware Technologies**

*Firmware technology and virtual hardware infrastructure.*
