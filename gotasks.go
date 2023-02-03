package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

const csvFile = "/.gotasks"

var homeDir, err = os.UserHomeDir()
var csvPath = homeDir + filepath.FromSlash(csvFile)
var repository = TasksCsvRepository{Path: csvPath}
var transformer = Transformer{}

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
	{
		Name:   "stop",
		Usage:  "Stop tracking a given task",
		Action: Stop,
	},
	{
		Name:   "status",
		Usage:  "Give status of all tasks",
		Action: Status,
	},
	{
		Name:   "clear",
		Usage:  "Clear all tasks",
		Action: Clear,
	},
	{
		Name:   "list",
		Usage:  "List all tasks",
		Action: List,
	},
}

func List(context *cli.Context) error {
	var err error
	transformer.LoadedTasks, err = repository.load()
	if err != nil {
		return err
	}

	for _, task := range transformer.Transform() {
		fmt.Println(task)
	}
	return nil
}

func Clear(context *cli.Context) error {
	err := repository.clear()
	if err == nil {
		fmt.Println("Cleared all tasks")
	}
	return err
}

func Status(context *cli.Context) error {
	identifer := context.Args().First()
	if !IsValidIdentifier(identifer) {
		return invalidIdentifier(identifer)
	}
	tasks, err := repository.load()
	if err != nil {
		return err
	}
	transformer.LoadedTasks = tasks.getByIdentifier(identifer)
	fmt.Println(transformer.Transform()[identifer])
	return nil
}

func Stop(context *cli.Context) error {
	identifier := context.Args().First()
	if !IsValidIdentifier(identifier) {
		return invalidIdentifier(identifier)
	}

	err := repository.save(Task{Identifier: identifier, Action: "stop", At: time.Now().Format(time.RFC3339)})

	if err == nil {
		fmt.Println("Stopped tracking task: " + identifier)
	}

	return err
}

func Start(c *cli.Context) error {
	identifier := c.Args().First()
	if !IsValidIdentifier(identifier) {
		return invalidIdentifier(identifier)
	}

	err := repository.save(Task{Identifier: identifier, Action: "start", At: time.Now().Format(time.RFC3339)})

	if err == nil {
		fmt.Println("Started tracking task ", identifier)
	}
	return err
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
	checkForInitialCSVFile()
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

func checkForInitialCSVFile() {
	if _, err := os.Stat(csvPath); os.IsNotExist(err) {
		os.Create(csvPath)
	}
}
