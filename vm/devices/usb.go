package devices

type USBDevice struct {
	Port    uint8
	NameID  string
	Enabled bool
}

func NewUSBDevice(port uint8, name string) *USBDevice {
	return &USBDevice{
		Port:   port,
		NameID: name,
	}
}

func (d *USBDevice) Name() string {
	return d.NameID
}

func (d *USBDevice) Reset() {
	d.Enabled = false
}

func (d *USBDevice) Start() error {
	d.Enabled = true
	return nil
}

func (d *USBDevice) Stop() error {
	d.Enabled = false
	return nil
}
