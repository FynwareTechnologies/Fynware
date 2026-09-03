package network

type Network struct {
	Enabled bool
}

func New(enabled bool) *Network {
	return &Network{
		Enabled: enabled,
	}
}

func (n *Network) Start() {
	n.Enabled = true
}

func (n *Network) Stop() {
	n.Enabled = false
}
