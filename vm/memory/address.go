package memory

type Address uint64

func (a Address) Offset(offset uint64) Address {
	return Address(uint64(a) + offset)
}
