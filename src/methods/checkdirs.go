package methods

import (
	"bufio"
	"filetools/src/util"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Checkdirs(dir string) {
	util.Info("The following is a list of all paths that will be removed:")

	time.Sleep(1 * time.Second)

	var paths []string
	scan(&paths, util.CWD(), dir)

	util.Warn("Do you want to delete these? THIS CANNOT BE UNDONE. (y/n, case-sensitive)")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	if strings.Contains(strings.ToLower(text), "y") {
		deldirs(&paths)
	} else {
		util.Fatal("User's input didn't contain y. Presuming they didn't want this. Aborting...");
		os.Exit(0)
	}
}

func deldirs(paths *[]string) {
	util.Warn("I sure hope you know what you're doing.\n")
	time.Sleep(1 * time.Second)

	for _, path := range *paths {
		util.Warn(fmt.Sprintf("Deleting directory %v", path))
		err := os.RemoveAll(path)
		if err != nil {
			util.Fatal(fmt.Sprintf("An error occured while trying to delete directory %v", path))
			os.Exit(1)
			return
		}
	}
}

func scan(paths *[]string, dir string, check string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name() == check {
				fmt.Printf(path + " ")
				*paths = append(*paths, path)
				return filepath.SkipDir
			}
		}
		return nil
	})

	if err != nil {
		util.Fatal("Some weird random thing went wrong while walking a directory")
		panic(err)
	}

	fmt.Println()
}
