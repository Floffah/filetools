package util

import "os"

func CWD() string {
	dir, err := os.Getwd()
	if err != nil {
		Fatal("Could not get the current working directory.")
		panic(err)
	}
	return dir
}