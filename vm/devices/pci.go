package devices

type PCIDevice struct {
	VendorID uint16
	DeviceID uint16
	Class    uint8
	Subclass uint8
	Bus      uint8
	Slot     uint8
	Function uint8
	Enabled  bool
}

func NewPCIDevice(vendorID, deviceID uint16, class, subclass uint8) *PCIDevice {
	return &PCIDevice{
		VendorID: vendorID,
		DeviceID: deviceID,
		Class:    class,
		Subclass: subclass,
	}
}

func (d *PCIDevice) Name() string {
	return "PCI Device"
}

func (d *PCIDevice) Reset() {
	d.Enabled = false
}

func (d *PCIDevice) Start() error {
	d.Enabled = true
	return nil
}

func (d *PCIDevice) Stop() error {
	d.Enabled = false
	return nil
}
