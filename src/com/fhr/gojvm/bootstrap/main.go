package main

import (
	"fmt"
	"com/fhr/gojvm/jconsole"
	"com/fhr/gojvm/core"
)

func main() {
	cmd := jconsole.ParseCmd()
	if (cmd.VersionFlag) {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		jconsole.PrintUsage()
	} else {
		core.StartJvm(cmd)
	}
}
