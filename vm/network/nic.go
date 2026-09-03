package network

type NIC struct {
	Name    string
	MAC     [6]byte
	Enabled bool
}

func NewNIC(name string, mac [6]byte) *NIC {
	return &NIC{
		Name: name,
		MAC:  mac,
	}
}

func (n *NIC) Start() {
	n.Enabled = true
}

func (n *NIC) Stop() {
	n.Enabled = false
}

func (n *NIC) Send(data []byte) int {
	if !n.Enabled {
		return 0
	}

	return len(data)
}
