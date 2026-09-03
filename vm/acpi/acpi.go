package acpi

type ACPI struct {
	Enabled bool
	Tables  *Tables
}

func New(enabled bool) *ACPI {
	return &ACPI{
		Enabled: enabled,
		Tables:  NewTables(),
	}
}

func (a *ACPI) Start() {
	a.Enabled = true
}

func (a *ACPI) Stop() {
	a.Enabled = false
}
