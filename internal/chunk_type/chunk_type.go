package chunktype

import "pngme/utils"

type ChunkType struct {
	Ancillary byte
	Private   byte
	Reserved  byte
	CopySafe  byte
}

func ChunkTypeFromStr(text string) *ChunkType {
	return &ChunkType{
		Ancillary: byte(text[0]),
		Private:   byte(text[1]),
		Reserved:  byte(text[2]),
		CopySafe:  byte(text[3]),
	}
}

func ChunkTypeFromBytes(bytes []byte) *ChunkType {
	return &ChunkType{
		Ancillary: bytes[0],
		Private:   bytes[1],
		Reserved:  bytes[2],
		CopySafe:  bytes[3],
	}
}

func getNthBit(b byte, nth int) byte {
	mask := 1 << nth
	return b & byte(mask)
}

func (c *ChunkType) GetTypeStr() string {
	return string(c.GetBytes())
}

func (c *ChunkType) GetBytes() []byte {
	return []byte{c.Ancillary, c.Private, c.Reserved, c.CopySafe}
}

func (c *ChunkType) IsValid() bool {
	return utils.CheckASCIIChar(c.GetBytes()) && c.IsReservedValid()
}

func (c *ChunkType) IsCritical() bool {
	return getNthBit(c.Ancillary, 5) == 0
}

func (c *ChunkType) IsPrivate() bool {
	return getNthBit(c.Private, 5) != 0
}

func (c *ChunkType) IsReservedValid() bool {
	return getNthBit(c.Reserved, 5) == 0
}

func (c *ChunkType) IsSafeToCopy() bool {
	return getNthBit(c.CopySafe, 5) != 0
}
