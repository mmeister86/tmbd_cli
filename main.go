package main

import "github.com/mmeister86/tmbd_cli/cmd"

// Version wird beim Build durch ldflags gesetzt
var Version = "1.0.2"

func main() {
	cmd.Execute()
}
