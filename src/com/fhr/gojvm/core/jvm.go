package core

import (
	"com/fhr/gojvm/jconsole"
	"fmt"
)

func StartJvm(cmd *jconsole.Cmd)  {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.CpOption, cmd.Class, cmd.Args)
}
