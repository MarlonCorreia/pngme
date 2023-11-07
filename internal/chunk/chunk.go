package chunk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	chunktype "pngme/internal/chunk_type"
)

const (
	IEND string = "IEND"
)

type Chunk struct {
	Len  uint32
	Type *chunktype.ChunkType
	Data []byte
	CRC  uint32
}

func getCRCChecksum(chunkType []byte, data []byte) uint32 {
	c := append(chunkType, data...)
	return crc32.ChecksumIEEE(c)
}

func NewChunk(data []byte, t *chunktype.ChunkType) *Chunk {
	crc := getCRCChecksum(t.GetBytes(), data)
	return &Chunk{
		Type: t,
		Len:  uint32(len(data)),
		Data: data,
		CRC:  crc,
	}
}

func (c *Chunk) GetFullChunkBytes() ([]byte, error) {
	buff := new(bytes.Buffer)
	buff.Grow(int(c.GetLength()) + 10)

	if err := binary.Write(buff, binary.BigEndian, c.GetLength()); err != nil {
		return nil, err
	}
	buff.Write(c.GetType().GetBytes())
	buff.Write(c.GetData())
	if err := binary.Write(buff, binary.BigEndian, c.GetCRC()); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

func NewChunkFromBytes(chunkData []byte) (*Chunk, uint32) {
	dataLen := binary.BigEndian.Uint32(chunkData[0:4])
	chunkTypeBytes := chunkData[4:8]
	data := chunkData[8 : 8+dataLen]
	crc := chunkData[8+dataLen : 12+dataLen]
	ct := chunktype.ChunkTypeFromBytes(chunkTypeBytes)

	return &Chunk{
		Len:  uint32(dataLen),
		Type: ct,
		Data: data,
		CRC:  binary.BigEndian.Uint32(crc),
	}, (12 + dataLen)
}

func (c *Chunk) PrintWholeChunk() {
	fmt.Println("data: ", string(c.Data))
	fmt.Println("len: ", c.GetLength())
	fmt.Println("crc: ", c.GetCRC())
	fmt.Println("type: ", c.GetType().GetTypeStr())
}

func (c *Chunk) GetCRC() uint32 {
	return c.CRC
}

func (c *Chunk) GetLength() uint32 {
	return c.Len
}

func (c *Chunk) GetType() *chunktype.ChunkType {
	return c.Type
}

func (c *Chunk) GetData() []byte {
	return c.Data
}

func (c *Chunk) GetDataStr() string {
	return string(c.Data)
}

func IENDChunk() *Chunk {
	ct := chunktype.ChunkTypeFromStr(IEND)
	return NewChunk([]byte{}, ct)
}
