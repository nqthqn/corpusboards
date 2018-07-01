package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var (
	errFlagParse = errors.New("failed to parse flags")
)

// mainConf ---------
type mainConf struct {
	fs      *flag.FlagSet
	verbose bool
}

func makeMainConf() mainConf {
	c := mainConf{
		fs: flag.NewFlagSet("main", flag.ContinueOnError),
	}

	return c
}

func (c *mainConf) attachFlags() {
	c.fs.BoolVar(&c.verbose, "v", c.verbose, "enable logging")
}

func (c *mainConf) normalize() error {
	return nil
}

// ngramConf ---------
type ngramConf struct {
	fs    *flag.FlagSet
	file  string
	n     int
	names bool
}

func makeNgramConf() ngramConf {
	c := ngramConf{
		fs:   flag.NewFlagSet("ngram", flag.ContinueOnError),
		file: "test_data",
	}

	return c
}

func (c *ngramConf) attachFlags() {
	c.fs.StringVar(&c.file, "f", c.file, "file to process")
	c.fs.BoolVar(&c.names, "names", c.names, "print unicode character names")
	c.fs.IntVar(&c.n, "n", c.n, "n-gram length (ex: 2 for bigram)")
}

func (c *ngramConf) normalize() error {
	if c.file == "" {
		return fmt.Errorf("file must not be empty string")
	}

	return nil
}

// Conf ... ---------
type Conf struct {
	cmd   string
	main  mainConf
	ngram ngramConf
}

func newConf() (*Conf, error) {
	c := &Conf{
		main:  makeMainConf(),
		ngram: makeNgramConf(),
	}

	return c, nil
}

func (c *Conf) parseFlags() error {
	c.main.attachFlags()
	c.ngram.attachFlags()

	if err := c.main.fs.Parse(os.Args[1:]); err != nil {
		return errFlagParse
	}

	if len(c.main.fs.Args()) == 0 {
		return nil
	}

	switch c.cmd = c.main.fs.Args()[0]; c.cmd {
	case c.ngram.fs.Name():
		if err := c.ngram.fs.Parse(nextArgs(os.Args, c.cmd)); err != nil {
			return errFlagParse
		}

		if err := c.ngram.normalize(); err != nil {
			return err
		}

	default:
		fmt.Fprintf(
			c.main.fs.Output(),
			"%q is not a valid subcommand, those available are: [%s]\n",
			c.cmd, c.ngram.fs.Name(),
		)

		return errFlagParse

	}

	return c.main.normalize()
}

func nextArgs(vals []string, val string) []string {
	for k, v := range vals {
		if v == val {
			return vals[k+1:]
		}
	}

	return vals
}
