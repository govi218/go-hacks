func (bits *BitList) REPLACE(index int, bit byte) *BitList {
	if bits.CheckIndex(index) != nil {
		// panic?
		panic(errors.New("Invalid index"))
	}

	byteIndex := uint32(index >> 3)
	byteOffset := uint32(index % 8)

	oldByte := &bits.ByteArray[byteIndex]
	*oldByte &= ^(1 << byteOffset)

	return bits
}

