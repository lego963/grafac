package main // import "github.com/lego963/grafac"
import (
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	args := os.Args[1:]

	if Commands == nil {
		initCommands()
	}

	// Build the CLI so far, we do this so we can query the subcommand.
	_ = &cli.CLI{
		Args:     args,
		Commands: Commands,
		//HelpFunc:   helpFunc,
		HelpWriter: os.Stdout,
	}
}
