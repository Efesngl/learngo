package types

type Command struct {
	Run         func(args []string)
	Description string
}
