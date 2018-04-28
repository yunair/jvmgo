package main

import (
	"fmt"
	"os"
	"flag"
)

type Cmd struct {
	helperFlag  bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [ars....]\n", os.Args[0])
}

func ParseCmd() *Cmd{
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.helperFlag, "help", false, "print http message")
	flag.BoolVar(&cmd.helperFlag, "?", false, "print http message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0{
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}
