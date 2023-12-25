package png

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"pngme/internal/chunk"
	chunktype "pngme/internal/chunk_type"
)

var (
	pngHead []byte = []byte{137, 80, 78, 71, 13, 10, 26, 10}
)

type Png struct {
	chunks []chunk.Chunk
}

func (p *Png) Head() []byte {
	return pngHead
}

func (p *Png) Chunks() []chunk.Chunk {
	return p.chunks
}

func (p *Png) ChunksByType(chunkType string) []chunk.Chunk {
	chunks := make([]chunk.Chunk, 0)
	for _, chunk := range p.Chunks() {
		if chunk.GetType().GetTypeStr() == chunkType {
			chunks = append(chunks, chunk)
		}
	}

	return chunks
}

func (p *Png) AddChunk(c *chunk.Chunk) {
	p.chunks = append(p.chunks, *c)
}

func (p *Png) RemoveChunk(chunkTypeStr string) {
	chunks := make([]chunk.Chunk, 0)
	for _, c := range p.Chunks() {
		if c.GetType().GetTypeStr() != chunkTypeStr {
			chunks = append(chunks, c)
		}
	}
	p.chunks = chunks
}

func (p *Png) PngAsBytes() []byte {
	data := pngHead

	for _, chunk := range p.Chunks() {
		d, _ := chunk.GetFullChunkBytes()
		data = append(data, d...)
	}
	cf, _ := chunk.IENDChunk().GetFullChunkBytes()
	data = append(data, cf...)
	return data
}

func PngFromFile(file io.Reader) (*Png, error) {
	png := new(Png)
	head := make([]byte, 8)

	_, err := io.ReadFull(file, head)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(pngHead, head) {
		return nil, fmt.Errorf("invalid png head")
	}

	chunks := make([]chunk.Chunk, 0)
	for {
		buff := make([]byte, 4)

		_, err := io.ReadFull(file, buff)
		if err != nil {
			return nil, err
		}
		len := binary.BigEndian.Uint32(buff)

		_, err = io.ReadFull(file, buff)
		if err != nil {
			return nil, err
		}

		if string(buff) == chunk.IEND {
			break
		}

		data := make([]byte, len)

		_, err = io.ReadFull(file, data)
		if err != nil {
			return nil, err
		}

		chunkType := chunktype.ChunkTypeFromBytes(buff)
		chunk := chunk.NewChunk(data, chunkType)

		// read 4 bytes to skip crc
		_, err = io.ReadFull(file, buff)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, *chunk)
	}

	png.chunks = chunks

	return png, nil
}

func PngFromChunks(chunks []chunk.Chunk) *Png {
	return &Png{
		chunks: chunks,
	}
}
