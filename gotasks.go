package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"regexp"
)

var commands = []*cli.Command{
	{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "show help",
		Action: func(c *cli.Context) error {
			fmt.Println("Commands:")
			return nil
		},
	},
	{
		Name:   "start",
		Usage:  "Start tracking a given task",
		Action: Start,
	},
}

func Start(c *cli.Context) error {
	identifier := c.Args().First()
	if !IsValidIdentifier(identifier) {
		return invalidIdentifier(identifier)
	}
	fmt.Println("Starting Task", identifier)
	return nil
}

func invalidIdentifier(identifier string) error {
	return fmt.Errorf("invalid identifier: %s", identifier)
}

func IsValidIdentifier(identifier string) bool {
	const alphanumericRegex = "^[a-zA-Z-1-9_-]+$"
	re := regexp.MustCompile(alphanumericRegex)
	return len(identifier) > 0 && re.MatchString(identifier)
}

func main() {
	app := cli.NewApp()
	app.Name = "Gotasks"
	app.Usage = "CLI timetracker for your tasks."
	app.Version = "0.0.1"
	app.Commands = commands
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
