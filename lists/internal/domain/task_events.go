package domain

type TaskAdded struct {
	Task *Task
}

func (TaskAdded) EventName() string { return "lists.TaskAdded" }

type TaskRemoved struct {
	Task *Task
}

func (TaskRemoved) EventName() string { return "lists.TaskRemoved" }

type TaskCompleted struct {
	Task *Task
}

func (TaskCompleted) EventName() string { return "lists.TaskCompleted" }

type TaskUncompleted struct {
	Task *Task
}

func (TaskUncompleted) EventName() string { return "lists.TaskUncompleted" }
