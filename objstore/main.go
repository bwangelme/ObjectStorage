package main

import (
	"flag"

	"github.com/bwangelme/ObjectStorage/api"
	"github.com/bwangelme/ObjectStorage/conf"
	"github.com/bwangelme/ObjectStorage/data"
)

func main() {
	flag.Parse()

	if conf.IsDataNode {
		data.Main()
	} else {
		api.Main()
	}
}
