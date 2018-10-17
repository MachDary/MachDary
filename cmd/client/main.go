package main

import (
	"runtime"

	cmd "github.com/MachDary/MachDary/cmd/client/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
