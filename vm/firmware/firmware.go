package firmware

type Firmware struct {
	Name    string
	Version string
	Image   []byte
	Loaded  bool
}

func New(name, version string, image []byte) *Firmware {
	return &Firmware{
		Name:    name,
		Version: version,
		Image:   image,
	}
}

func (f *Firmware) Load() {
	f.Loaded = true
}

func (f *Firmware) Unload() {
	f.Loaded = false
}
