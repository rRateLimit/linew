package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/rRateLimit/linew/internal/config"
	"github.com/rRateLimit/linew/internal/wrap"
)

func main() {
	cfg := parseFlags()

	var reader io.Reader
	if len(flag.Args()) > 0 {
		file, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	var writer io.Writer
	if cfg.Output != "" {
		file, err := os.Create(cfg.Output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	if err := processText(reader, writer, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing text: %v\n", err)
		os.Exit(1)
	}
}

func parseFlags() *config.Config {
	cfg := &config.Config{}

	flag.IntVar(&cfg.Width, "width", 80, "Maximum width for line wrapping")
	flag.IntVar(&cfg.Width, "w", 80, "Maximum width for line wrapping (shorthand)")
	flag.BoolVar(&cfg.PreserveIndent, "indent", true, "Preserve indentation")
	flag.BoolVar(&cfg.PreserveIndent, "i", true, "Preserve indentation (shorthand)")
	flag.StringVar(&cfg.Output, "output", "", "Output file (default: stdout)")
	flag.StringVar(&cfg.Output, "o", "", "Output file (shorthand)")

	noIndent := flag.Bool("no-indent", false, "Do not preserve indentation")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [file]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "A line wrapping tool for text files.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s input.txt                  # Wrap lines at 80 characters\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -w 100 input.txt           # Wrap lines at 100 characters\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s --no-indent input.txt      # Wrap without preserving indentation\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  echo \"text\" | %s              # Process from stdin\n", os.Args[0])
	}

	flag.Parse()

	if *noIndent {
		cfg.PreserveIndent = false
	}

	return cfg
}

func processText(reader io.Reader, writer io.Writer, cfg *config.Config) error {
	scanner := bufio.NewScanner(reader)
	w := wrap.New(cfg)

	for scanner.Scan() {
		line := scanner.Text()
		wrapped := w.WrapLine(line)
		
		for _, wl := range wrapped {
			if _, err := fmt.Fprintln(writer, wl); err != nil {
				return err
			}
		}
	}

	return scanner.Err()
}