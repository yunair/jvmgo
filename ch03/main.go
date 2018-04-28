package main

import (
	"flag"
	"fmt"
	"github.com/yunair/jvmgo/ch02/classpath"
	"strings"
	"github.com/yunair/jvmgo/ch03/classfile"
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
	cf := loadClass(className, cp)
	fmt.Println(cmd.Class)
	printClassInfo(cf)
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf ,err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}