package main

import (
	"errors"
	"strings"
	// "flag"
	"fmt"
	"os"
	// "strings"
)

func main() {
	src, err := readInput()
	
	if err != nil {
		fail(err)
	}
	if (src == "") {
		fmt.Println(0)	
	} else {
		fmt.Println(len(strings.Split(strings.Trim(src," ")," ")))	
	}
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	
	if len(os.Args) < 2 {
		return  src, errors.New("missing argument")
	}

	src = os.Args[1]
	
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
