package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("missing argument")
		os.Exit(1)
	}

	regEx, err := regexp.Compile("^.+\\.(" + os.Args[2] + ")$")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err == nil && regEx.MatchString(info.Name()) {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
