package domain

type ListCreated struct {
	List *List
}

func (ListCreated) EventName() string { return "lists.ListCreated" }
