package main

type Tasks struct {
	Items []Task
}

func (tasks *Tasks) addItem(item Task) []Task {
	tasks.Items = append(tasks.Items, item)
	return tasks.Items
}

func (tasks *Tasks) getByIdentifier(identifer string) Tasks {
	tasksWithIdentifier := Tasks{}
	for _, task := range tasks.Items {
		if task.getIdentifier() != identifer {
			continue
		}
		tasksWithIdentifier.addItem(task)
	}
	return tasksWithIdentifier
}

type Task struct {
	Identifier string
	Action     string
	At         string
}

func (task Task) toArrayString() []string {
	return []string{task.Identifier, task.Action, task.At}
}

func (task Task) getIdentifier() string {
	return task.Identifier
}

func (task Task) getAction() string {
	return task.Action
}

func (task Task) getAt() string {
	return task.At
}
