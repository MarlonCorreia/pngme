package chunk_test

import (
	"bytes"
	"pngme/internal/chunk"
	chunktype "pngme/internal/chunk_type"
	"testing"
)

var (
	chunkTypeStr   string               = "RuSt"
	chunkType      *chunktype.ChunkType = chunktype.ChunkTypeFromStr(chunkTypeStr)
	data           []byte               = []byte("This is where your secret message will be!")
	crc            uint32               = 2882656334
	length         uint32               = 42
	chunkFullBytes []byte               = []byte{
		0, 0, 0, 42, 82, 117, 83, 116, 84, 104, 105, 115,
		32, 105, 115, 32, 119, 104, 101, 114, 101, 32, 121,
		111, 117, 114, 32, 115, 101, 99, 114, 101, 116, 32,
		109, 101, 115, 115, 97, 103, 101, 32, 119, 105, 108,
		108, 32, 98, 101, 33, 171, 209, 216, 78}
)

func getChunk() *chunk.Chunk {
	return chunk.NewChunk(data, chunkType)
}

func TestChunkLength(t *testing.T) {
	c := getChunk()

	if c.GetLength() != length {
		t.Errorf("Expected %d got %d", length, c.GetLength())
	}
}

func TestChunkCRC(t *testing.T) {
	c := getChunk()

	if c.GetCRC() != crc {
		t.Errorf("Expected %d got %d", crc, c.GetCRC())
	}
}

func TestChunkType(t *testing.T) {
	c := getChunk()
	got := c.GetType().GetTypeStr()
	if got != chunkTypeStr {
		t.Errorf("Expected %s got %s", chunkTypeStr, got)
	}
}

func TestChunkDataStr(t *testing.T) {
	c := getChunk()

	if c.GetDataStr() != string(data) {
		t.Errorf("Expected %s got %s", string(data), c.GetDataStr())
	}
}

func TestChunkFullBytes(t *testing.T) {
	c := getChunk()
	chunkBytes, err := c.GetFullChunkBytes()
	if err != nil {
		t.Error("Error while getting chunk bytes")
	}

	if !bytes.Equal(chunkFullBytes, chunkBytes) {
		t.Errorf("Expected %d got %d", chunkFullBytes, chunkBytes)
	}

}
