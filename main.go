package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func usage() {
	argv0 := path.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "usage: %s [-dryrun|-n] [-config|-c <filename>][-c ...]", argv0)
	fmt.Fprintln(os.Stderr, "")
	flag.PrintDefaults() // defaults to print to stderr
	os.Exit(2)
}

func main() {
	var (
		dryrun bool
		config Config
	)

	flag.Var(&config, "config", "Configuration JSON (multiples appended)")
	flag.Var(&config, "c", "Configuration JSON (multiples appended)")

	flag.BoolVar(&dryrun, "dryrun", dryrun, "Parse, show, do not run")
	flag.BoolVar(&dryrun, "n", dryrun, "Parse, show, do not run")

	flag.Usage = usage
	flag.Parse()

	if dryrun {
		fmt.Fprintln(os.Stderr, "DRYRUN    DRYRUN     DRYRUN")
		fmt.Fprintln(os.Stderr, config.String())
	} else {
		for _, v := range config {
			// errors are due to os/exec.Cmd(...).Run() returns
			if err := MoveClient.Copy(
				os.ExpandEnv(v.Destination),
				os.ExpandEnv(v.Source),
			); err != nil {
				fmt.Fprintf(os.Stderr, "Error moving from \"%s\" to \"%s\": %v\n", v.Source, v.Destination, err)
			}
		}
	}

}
