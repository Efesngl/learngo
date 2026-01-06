package types

type TodoItem struct {
	Id     int
	Title  string
	IsDone bool
}

type Command struct {
	Name        string
	Fn          string
	Description string
	Hint        string
	Run         func(args []string)
}
