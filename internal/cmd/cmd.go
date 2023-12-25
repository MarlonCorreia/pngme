package cmd

import (
	"flag"
	"fmt"
	"os"
	"pngme/internal/controllers"
)

var(
    encode = "encode"
    decode = "decode"
    remove = "remove"

)

func InitCmd() {
    encodeCmd := flag.NewFlagSet(encode, flag.ExitOnError)
    encFile := encodeCmd.String("file", "", "png file")
    encMsg := encodeCmd.String("message", "", "message to encode")
    encHead := encodeCmd.String("head", "", "head of chunk type")

    decodeCmd := flag.NewFlagSet(decode, flag.ExitOnError)
    decFile := decodeCmd.String("file", "", "png file")
    decHead := decodeCmd.String("head", "", "head of chunk type")
    
    removeCmd := flag.NewFlagSet(remove, flag.ExitOnError)
    remFile := removeCmd.String("file", "", "png file")
    remHead := removeCmd.String("head", "", "head of chunk type")
    
    if len(os.Args) == 1 {
        fmt.Println("Please provide a command")
        os.Exit(0)
    }

    switch os.Args[1] {
    case encode: 
        encodeCmd.Parse(os.Args[2:])
        err := controllers.EncodeMsg(encFile, encMsg, encHead)
        if err != nil {
            fmt.Println("Unable to encode message to file")
            os.Exit(0)
        }


    case decode:
        decodeCmd.Parse(os.Args[2:])
        err := controllers.DecodeMsg(decFile, decHead) 
        if err != nil {
            fmt.Println("Unable to decode message from file")
            os.Exit(0)
        }

    case remove:
        removeCmd.Parse(os.Args[2:])
        err := controllers.RemoveMsg(remFile, remHead)
        if err != nil {
            fmt.Println("Unable to remove message from file")
            os.Exit(0)
        }

    default:
        fmt.Println("No valid command provided")
    }

}
