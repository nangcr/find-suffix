package main

import (
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		println("missing argument")
		os.Exit(1)
	}

	regEx, err := regexp.Compile("^.+\\.(" + os.Args[1] + ")$")
	if err != nil {
		println(err)
		os.Exit(1)
	}

	err = filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err == nil && regEx.MatchString(info.Name()) {
			println(path)
		}
		return nil
	})
	if err != nil {
		println(err)
		os.Exit(1)
	}

}
