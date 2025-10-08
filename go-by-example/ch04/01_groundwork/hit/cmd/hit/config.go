package main

import (
	"fmt"
	"strconv"
	"strings"
)

// parseFunc defines a function type that parses a string
// and returns an error if parsing fails.
type parseFunc func(string) error

// stringVar returns a closure that assigns the given string `s`
// to the string variable pointed to by `p`.
// This is a higher-order function: it takes a pointer and
// returns a function that “remembers” that pointer.
func stringVar(p *string) parseFunc {
	return func(s string) error {
		*p = s
		return nil
	}
}

// intVar returns a closure that parses the string `s` into an int
// and assigns it to the int variable pointed to by `p`.
// If strconv.Atoi fails, the error is returned.
func intVar(p *int) parseFunc {
	return func(s string) error {
		var err error
		*p, err = strconv.Atoi(s)
		return err
	}
}

type config struct {
	url string
	n   int
	c   int
	rps int
}

func parseArgs(c *config, args []string) error {
	flagSet := map[string]parseFunc{
		// Use & to get pointers to struct fields because
		// field access on a pointer (c.url) is auto-dereferenced in Go.
		"url": stringVar(&c.url),
		"n":   intVar(&c.n),
		"c":   intVar(&c.c),
		"rps": intVar(&c.rps),
	}

	for _, arg := range args {
		name, val, _ := strings.Cut(arg, "=")
		name = strings.TrimPrefix(name, "-")

		setVar, ok := flagSet[name]
		if !ok {
			return fmt.Errorf("flag provided but not defined: -%s", name)
		}
		if err := setVar(val); err != nil {
			return fmt.Errorf("invalid value %q for flag -%s: %w", val, name, err)
		}
	}
	return nil
}
