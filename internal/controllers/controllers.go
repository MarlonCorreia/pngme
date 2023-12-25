package controllers

import (
	"fmt"
	"os"
	"pngme/internal/chunk"
	chunktype "pngme/internal/chunk_type"
	"pngme/internal/png"
	"pngme/utils"
)

var(
    baseOutput = "output.png"
)

func EncodeMsg(file *string, msg *string, head *string) error {
    f, err := os.Open(*file)
    if err != nil {
        return err
    }

    png, err := png.PngFromFile(f)
    if err != nil {
        return err
    }

    ct := chunktype.ChunkTypeFromStr(*head)
    c := chunk.NewChunk([]byte(*msg), ct)
    png.AddChunk(c)
    utils.WriteNewFile(baseOutput, png.PngAsBytes())
    
    return nil 
}

func DecodeMsg(file *string, head *string) error {
    f, err := os.Open(*file)
    if err != nil {
        return err
    }

    png, err := png.PngFromFile(f)
    if err != nil {
        return err
    }
    
    chunks := png.ChunksByType(*head)
    if len(chunks) == 0 {
        fmt.Println("No message found")
    } else {
        for _, c := range chunks {
            fmt.Println( c.GetDataStr()) 
        }

    }
    return nil
}


func RemoveMsg(file *string, head *string) error {
    f, err := os.Open(*file)
    if err != nil {
        return err
    }
    
    png, err := png.PngFromFile(f)
    if err != nil {
        return err
    }
 
    png.RemoveChunk(*head)
    utils.WriteNewFile(baseOutput, png.PngAsBytes())
    return nil
}
