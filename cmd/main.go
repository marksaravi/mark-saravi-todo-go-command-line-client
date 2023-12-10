package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"gitgub.com/marksaravi/mark-saravi-todo-go-command-line-client/todos"
	"github.com/lpernett/godotenv"
)

type todosReporter interface {
	GetTodos()
	ToDosReport()
}

var commandsMappings = map[string]bool{
	"even":   true,
	"odd":    true,
	"custom": true,
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	reporter := TodoCommandFactory()
	st := time.Now()
	reporter.GetTodos()
	dur := time.Since(st)
	reporter.ToDosReport()
	fmt.Printf("Dur(ms): %d\n", dur.Milliseconds())
}

func TodoCommandFactory() todosReporter {
	methodPtr := flag.String("method", "even", "a string")
	fromPtr := flag.String("from", "2", "a string")
	nPtr := flag.String("n", "20", "a string")
	flag.Parse()
	method := "even"
	from := 2
	n := 20
	if v, err := strconv.Atoi(*fromPtr); err == nil {
		from = v
	}
	if v, err := strconv.Atoi(*nPtr); err == nil {
		n = v
	}
	if *methodPtr == "odd" {
		method = "odd"
	}

	switch method {
	case "even":
		return todos.NewEvenTODOs(from, n)
	case "odd":
		return todos.NewOddTODOs(from, n)
	default:
		return todos.NewEvenTODOs(from, n)
	}
}
