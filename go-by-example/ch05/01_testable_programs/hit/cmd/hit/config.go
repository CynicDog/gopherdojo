package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

type config struct {
	url string
	n   int
	c   int
	rps int
}

func parseArgs(c *config, args []string, stderr io.Writer) error {
	fs := flag.NewFlagSet(
		"hit",
		flag.ContinueOnError,
	)
	fs.SetOutput(stderr)

	fs.Usage = func() {
		fmt.Fprintf(fs.Output(), "usage: %s [options] url\n", fs.Name())
		fs.PrintDefaults()
	}

	fs.Var(asPositiveIntValue(&c.n), "n", "Number of requests")
	fs.Var(asPositiveIntValue(&c.c), "c", "Concurrency level")
	fs.Var(asPositiveIntValue(&c.rps), "rps", "Requests per second")

	if err := fs.Parse(args); err != nil {
		return err
	}
	c.url = fs.Arg(0)

	if err := validateArgs(c); err != nil {
		fmt.Fprintln(fs.Output(), err)
		fs.Usage()
		return err
	}

	return nil
}

// positiveIntValue is a named type whose underlying type is int.
// It implements the flag.Value interface to allow defining a flag
// that only accepts positive integer values.
type positiveIntValue int

// asPositiveIntValue converts an *int to a *positiveIntValue.
// It does not allocate new memory; both pointers refer to the same variable.
func asPositiveIntValue(p *int) *positiveIntValue {
	return (*positiveIntValue)(p)
}

// String returns the string representation of the positiveIntValue.
// It satisfies the String method required by the flag.Value interface.
func (n *positiveIntValue) String() string {
	return strconv.Itoa(int(*n))
}

// Set parses the provided string as an integer and assigns it to n.
// The value must be greater than zero, otherwise Set returns an error.
// This method satisfies the Set method required by the flag.Value interface.
func (n *positiveIntValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, strconv.IntSize)
	if err != nil {
		return err
	}
	if v <= 0 {
		return errors.New("should be greater than zero")
	}
	*n = positiveIntValue(v)
	return nil
}

func validateArgs(c *config) error {
	u, err := url.Parse(c.url)
	if err != nil {
		return fmt.Errorf("invalid value %q for url: %w", c.url, err)
	}
	if c.url == "" || u.Host == "" || u.Scheme == "" {
		return fmt.Errorf("invalud value %q for url: requires a valid url", c.url)
	}
	if c.n < c.c {
		return fmt.Errorf("invalud value %d for flag -n: should be greater than flag -c: %d", c.n, c.c)
	}
	return nil
}
