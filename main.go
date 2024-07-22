package main

import (
	"os"

	"github.com/Lentanta/FantasyScript/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
