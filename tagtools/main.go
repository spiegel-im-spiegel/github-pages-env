package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/goark/errs"
	"github.com/goark/struct2pflag"
	"github.com/spf13/pflag"

	"github-pages-env/tagtools/internal/tagslist"
	"github-pages-env/tagtools/internal/toptags"
	"github-pages-env/tagtools/internal/verify"
)

type jsonOutputError struct {
	Payload string
}

func (e *jsonOutputError) Error() string {
	return "verification failed"
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		var jerr *jsonOutputError
		if errors.As(err, &jerr) {
			fmt.Fprintln(os.Stderr, jerr.Payload)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		printUsage()
		return nil
	}

	subcmd := args[0]
	subArgs := args[1:]

	switch subcmd {
	case "tagslist":
		cfg := tagslist.DefaultConfig()
		fs := pflag.NewFlagSet("tagslist", pflag.ContinueOnError)
		struct2pflag.Bind(fs, &cfg)
		if err := fs.Parse(subArgs); err != nil {
			return errs.Wrap(err)
		}
		if err := tagslist.Run(cfg); err != nil {
			return errs.Wrap(err)
		}
		fmt.Println("Updated", cfg.Out)
		return nil
	case "toptags":
		cfg := toptags.DefaultConfig()
		fs := pflag.NewFlagSet("toptags", pflag.ContinueOnError)
		struct2pflag.Bind(fs, &cfg)
		if err := fs.Parse(subArgs); err != nil {
			return errs.Wrap(err)
		}
		if err := toptags.Run(cfg); err != nil {
			return errs.Wrap(err)
		}
		fmt.Println("Updated", cfg.Out)
		return nil
	case "all":
		if err := toptags.Run(toptags.DefaultConfig()); err != nil {
			return errs.Wrap(err)
		}
		if err := tagslist.Run(tagslist.DefaultConfig()); err != nil {
			return errs.Wrap(err)
		}
		fmt.Println("Updated data/toptags.json and .github/workflows/tagslist.csv")
		return nil
	case "verify":
		cfg := verify.DefaultConfig()
		fs := pflag.NewFlagSet("verify", pflag.ContinueOnError)
		struct2pflag.Bind(fs, &cfg)
		if err := fs.Parse(subArgs); err != nil {
			return errs.Wrap(err)
		}
		if err := verify.Run(cfg); err != nil {
			if cfg.Debug {
				payload, jerr := verify.JSONError(err, cfg)
				if jerr != nil {
					return errs.Wrap(jerr)
				}
				return &jsonOutputError{Payload: payload}
			}
			return errs.Wrap(err)
		}
		if cfg.Debug {
			payload, jerr := verify.JSONSuccess(cfg)
			if jerr != nil {
				return errs.Wrap(jerr)
			}
			fmt.Println(payload)
			return nil
		}
		fmt.Println("Verification passed")
		return nil
	case "help", "-h", "--help":
		printUsage()
		return nil
	default:
		return errs.New("unknown subcommand", errs.WithContext("subcommand", subcmd))
	}
}

func printUsage() {
	fmt.Println("tagtools: helper CLI for tagslist/toptags")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  tagtools tagslist [flags]")
	fmt.Println("  tagtools toptags [flags]")
	fmt.Println("  tagtools all")
	fmt.Println("  tagtools verify [flags]")
}
