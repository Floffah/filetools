package main

import (
	"filetools/src/util"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	util.Info("FileTools by Floffah v1.0.0\n")

	var dirname string
	flag.StringVar(&dirname, "Name", "", "-Name=dirToUnLink")

	flag.Parse()

	if dirname == "" {
		util.Fatal("Option `-Name` cannot be empty.")
		os.Exit(1)
	}

	util.Info(fmt.Sprintf("Searching recursively for directories named: %v", dirname))

	process(dirname)
}

func process(dir string) {
	util.Info("The following is a list of all paths that will be removed:")

	time.Sleep(1 * time.Second)

	var paths []string
	scan(&paths, filepath.Join(util.CWD(), dir), dir)
}

func scan(paths *[]string, dir string, check string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if !strings.HasSuffix(info.Name(), check) {
				defer scan(paths, path, check)
			} else {
				fmt.Printf(path)
				*paths = append(*paths, path)
			}
		}
		return nil
	})
	if err != nil {
		util.Fatal("Some weird random thing went wrong while walking a directory")
		panic(err)
	}
}

// 				err := syscall.Rmdir(filepath.Join(path))
//				if err != nil {
//					return err
//				}