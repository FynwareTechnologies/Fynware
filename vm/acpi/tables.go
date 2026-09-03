package acpi

type TableHeader struct {
	Signature [4]byte
	Length    uint32
	Revision  uint8
	Checksum  uint8
}

type Tables struct {
	RSDP *RSDP
	RSDT *RSDT
	FADT *FADT
	DSDT *DSDT
}

type RSDP struct {
	Signature [8]byte
	Revision  uint8
}

type RSDT struct {
	Header TableHeader
}

type FADT struct {
	Header TableHeader
}

type DSDT struct {
	Header TableHeader
	AML    []byte
}

func NewTables() *Tables {
	return &Tables{
		RSDP: &RSDP{
			Signature: [8]byte{'R', 'S', 'D', ' ', 'P', 'T', 'R', ' '},
			Revision:  2,
		},
		RSDT: &RSDT{
			Header: newHeader("RSDT"),
		},
		FADT: &FADT{
			Header: newHeader("FACP"),
		},
		DSDT: &DSDT{
			Header: newHeader("DSDT"),
			AML:    []byte{},
		},
	}
}

func newHeader(signature string) TableHeader {
	var header TableHeader
	copy(header.Signature[:], signature)
	return header
}
