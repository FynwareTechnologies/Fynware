package display

type GPU struct {
	Name    string
	Enabled bool
}

func NewGPU(name string) *GPU {
	return &GPU{
		Name: name,
	}
}

func (g *GPU) Start() {
	g.Enabled = true
}

func (g *GPU) Stop() {
	g.Enabled = false
}
