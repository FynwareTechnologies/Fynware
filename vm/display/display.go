package display

type Display struct {
	Width   uint32
	Height  uint32
	Enabled bool
}

func New(width, height uint32) *Display {
	return &Display{
		Width:   width,
		Height:  height,
		Enabled: true,
	}
}

func (d *Display) Start() {
	d.Enabled = true
}

func (d *Display) Stop() {
	d.Enabled = false
}
