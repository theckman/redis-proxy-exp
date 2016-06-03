package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/jessevdk/go-flags"
)

type binArgs struct {
	Port      uint16 `short:"p" long:"port" default:"1620" description:"port to bind to for HTTP interface"`
	RedisHost string `long:"redis-host" required:"true" description:"IP or hostname or Redis instance (required)"`
	RedisPort uint16 `long:"redis-port" default:"6379" description:"port number where Redis is listening"`
	Version   bool   `short:"V" long:"version" description:"print versions string and exit"`
}

func (a *binArgs) parse(args []string) (string, error) {
	if args == nil {
		args = os.Args
	}

	p := flags.NewParser(a, flags.HelpFlag|flags.PassDoubleDash)

	_, err := p.ParseArgs(args[1:])

	// determine if there was a parsing error
	// unfortunately, help message is returned as an error
	if err != nil {
		// determine whether this was a help message by doing a type
		// assertion of err to *flags.Error and check the error type
		// if it was a help message, do not return an error
		if errType, ok := err.(*flags.Error); ok && errType.Type == flags.ErrHelp {
			return err.Error(), nil
		}

		return "", err
	}

	if a.Version {
		out := fmt.Sprintf(
			"fc-interview v%s built with %s\n",
			fciVersion, runtime.Version(),
		)

		return out, nil
	}

	return "", nil
}
