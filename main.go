package main

import (
	"flag"
	"fmt"
	"github.com/roblaszczak/go-cleanarch/cleanarch"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	var path string

	if len(flag.Args()) > 1 {
		path = flag.Args()[1]
	} else {
		var err error
		path, err = filepath.Abs(flag.Arg(0))
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("[cleanarch] checking %s\n", path)

	validator := cleanarch.NewValidator()
	isValid, errors, err := validator.Validate(path)
	if err != nil {
		panic(err)
	}

	if !isValid {
		for _, err := range errors {
			fmt.Println(err.Error())
		}

		fmt.Println("Uncle Bob is not happy.")
		os.Exit(1)
	}

	os.Exit(0)
}
