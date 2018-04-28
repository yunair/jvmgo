package main

import (
	"fmt"
	"os"
)

type Cmd struct {
	HelperFlag  bool
	VersionFlag bool
	CpOption    string
	XjreOption  string
	Class       string
	Args        []string
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [ars....]\n", os.Args[0])
}
