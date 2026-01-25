package main

import "tmdb-cli/cmd"

// Version wird beim Build durch ldflags gesetzt
var Version = "1.0.1"

func main() {
	cmd.Execute()
}
