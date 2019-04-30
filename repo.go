package main

var (
	todos     Todos
	currentID int
)

func init() {
	RepoCreateTodo(Todo{Name: "Write presentation"})
	RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoCreateTodo(t Todo) Todo {
	currentID += 1
	t.ID = currentID
	todos = append(todos, t)
	return t
}
