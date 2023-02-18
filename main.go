package main // import "github.com/lego963/grafac"

import (
	"os"

	"github.com/lego963/grafac/internal"
)

func init() {
	// this is a good place to patch SHA-1 support back into x509
	internal.PatchSha1()
}

func main() {
	os.Exit(command.Run(os.Args[1:]))
}
