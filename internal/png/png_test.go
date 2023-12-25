package png_test

import (
	"bytes"
	"pngme/internal/chunk"
	chunktype "pngme/internal/chunk_type"
	"pngme/internal/png"
	"pngme/utils"
	"testing"
)

var (
    pngExample []byte = utils.GetPNGExample()
)

func createChunkAndType(data string, typestr string) chunk.Chunk {
	ct := chunktype.ChunkTypeFromStr(typestr)
	return *chunk.NewChunk([]byte(data), ct)
}

func createChunks() []chunk.Chunk {
	return []chunk.Chunk{
		createChunkAndType("I am the first", "FrSt"),
		createChunkAndType("I am the middle", "miDl"),
		createChunkAndType("I am the last", "LASt"),
	}
}

func chunksPng() *png.Png {
	return png.PngFromChunks(createChunks())
}

func TestPngFromChunks(t *testing.T) {
	png := chunksPng()
	chunksLen := len(png.Chunks())

	if chunksLen != 3 {
		t.Errorf("Expected %d got %d", 3, chunksLen)
	}
}

func TestPngFromBytes(t *testing.T) {
    f2 := bytes.NewReader(pngExample) 
    _, err := png.PngFromFile(f2)
	if err != nil {
		t.Error("Failed creating png from file")
	}
}
