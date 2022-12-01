package main

import (
	mtool "github.com/surdeus/gomtool/src/multitool"
	"github.com/surdeus/gore/src/tool/bin2hex"
)

func main() {
	tools := mtool.Tools {
		"bin2hex" : mtool.Tool{bin2hex.Run, "convert binary stdin to hex"},
	}

	mtool.Main("gore", tools)
}

