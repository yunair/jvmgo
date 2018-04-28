package main

import (
	"flag"
	"fmt"
	"github.com/yunair/jvmgo/ch02/classpath"
	"strings"
)


func ParseCmd() *Cmd{
	cmd := &Cmd{}
	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelperFlag, "help", false, "print http message")
	flag.BoolVar(&cmd.HelperFlag, "?", false, "print http message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0{
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}
	return cmd
}

func main() {
	cmd := ParseCmd()
	if cmd.VersionFlag {
		fmt.Println("Version 0.0.1")
	} else if cmd.HelperFlag || cmd.Class == "" {
		PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Printf("classpath : %s class : %s args : %v\n",
		cmd.CpOption, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _ , err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.Class)
		return
	}

	fmt.Printf("class data:%v\n", classData)
}