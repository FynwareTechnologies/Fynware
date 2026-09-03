package platform

type Platform struct {
	Name     string
	Hardware *Hardware
}

func New(name string, hardware *Hardware) *Platform {
	return &Platform{
		Name:     name,
		Hardware: hardware,
	}
}
