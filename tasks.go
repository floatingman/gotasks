package main

type Tasks struct {
	Items []Task
}

type Task struct {
	Identifier string
	Action     string
	At         string
}

func (task Task) toArrayString() []string {
	return []string{task.Identifier, task.Action, task.At}
}
