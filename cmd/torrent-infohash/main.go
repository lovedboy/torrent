package main

import (
	"fmt"
	"log"

	"github.com/anacrolix/tagflag"

	"github.com/lovedboy/torrent/metainfo"
)

func main() {
	var args struct {
		tagflag.StartPos
		Files []string `arity:"+" type:"pos"`
	}
	tagflag.Parse(&args)
	for _, arg := range args.Files {
		mi, err := metainfo.LoadFromFile(arg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", mi.Info.Hash().HexString(), arg)
	}
}
