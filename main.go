package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/pksunkara/alpaca/alpaca"
	"os"
)

func main() {

	// Options for flags package
	var opts struct {
		Version bool `short:"v" long:"version" description:"Show version information"`

		Langs alpaca.LanguageOptions `group:"Language Options"`
	}

	// Build the parser
	parser := flags.NewParser(&opts, flags.Default)

	// Set usage string
	parser.Usage = "[options] <dir>"

	// Parse the arguments
	args, err := parser.Parse()

	if err != nil {
		os.Exit(0)
	}

	// Print version and exit
	if opts.Version {
		fmt.Println(alpaca.Version)
		os.Exit(0)
	}

	// If no argument is given
	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	alpaca.WriteLibraries(args[0], &opts.Langs)
}
