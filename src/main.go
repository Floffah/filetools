package main

import (
	"filetools/src/methods"
	"filetools/src/util"
	"flag"
	"fmt"
	"os"
)

func main() {
	util.Info("FileTools by Floffah v1.0.0\n")

	var method string
	var dirname string
	flag.StringVar(&method, "Method", "", "-Method=checkdir")
	flag.StringVar(&dirname, "Name", "", "-Name=dirToUnLink")

	flag.Parse()

	if method == "checkdir" {
		if dirname == "" {
			util.Fatal("Option `-Name` cannot be empty.")
			os.Exit(1)
		}

		util.Info(fmt.Sprintf("Searching recursively for directories named: %v", dirname))

		methods.Checkdirs(dirname)
	} else {
		util.Fatal("Flag -Method must be one of the following values: checkdir")
		os.Exit(1)
		return
	}

	util.Info("Done!")
}
