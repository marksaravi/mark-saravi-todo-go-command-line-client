package main

import (
	"flag"
	"os"

	"gitgub.com/marksaravi/mark-saravi-todo-go-command-line-client/todos"
)

type todosReporter interface {
	ToDosReport()
}

var commandsMappings = map[string]bool{
	"even":   true,
	"odd":    true,
	"custom": true,
}

func main() {
	arguments := os.Args[1:]
	var command string
	if len(arguments) == 0 {
		command = "even"
	} else if _, ok := commandsMappings[arguments[0]]; ok {
		command = arguments[0]
	} else {
		command = "even"
	}
	reporter := TodoCommandFactory(command)
	reporter.ToDosReport()
}

func TodoCommandFactory(command string) todosReporter {
	fromPtr := flag.Int("from", 2, "an int")
	nPtr := flag.Int("n", todos.DEFAULT_NUMBER_OF_TODOS, "an int")
	idsPtr := flag.String("range", "2,4,6,8,10", "a string")
	from := *fromPtr
	n := *nPtr
	ids := *idsPtr
	switch command {
	case "even":
		return todos.NewEvenTODOs(from, n)
	case "odd":
		return todos.NewOddTODOs(from, n)
	case "custom":
		return todos.NewCustomTODOs(ids)
	default:
		return todos.NewEvenTODOs(from, n)
	}
}
