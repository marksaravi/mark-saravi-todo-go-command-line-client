package main

import (
	"fmt"

	todoapi "gitgub.com/marksaravi/mark-saravi-todo-go-command-line-client/api/todo-api-client"
)

func main() {
	var ids []int
	for i := 2; i <= 20; i += 2 {
		ids = append(ids, i)
	}
	fmt.Println(ids)
	todoApiClient := todoapi.NewToDoApiClient("https://jsonplaceholder.typicode.com/todos")
	todoApiClient.GetTODOs(ids...)
}
