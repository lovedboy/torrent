package main

import (
	"fmt"
	"os"

	"github.com/anacrolix/tagflag"

	"github.com/lovedboy/torrent/metainfo"
)

func main() {
	tagflag.Parse(nil)

	mi, err := metainfo.Load(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading metainfo from stdin: %s", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", mi.Magnet().String())
}
