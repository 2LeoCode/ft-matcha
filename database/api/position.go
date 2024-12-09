package api

import "encoding/binary"

type Position struct {
	Lattitude int8
	Longitude int8
}

func (this *Position) Encode() uint16 {
	return binary.BigEndian.Uint16([]byte{byte(this.Lattitude), byte(this.Longitude)})
}

func (this *Position) Decode(encoded uint16) {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, encoded)
	this.Lattitude = int8(bytes[0])
	this.Longitude = int8(bytes[1])
}
