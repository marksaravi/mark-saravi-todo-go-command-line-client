package todos

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	api "gitgub.com/marksaravi/mark-saravi-todo-go-command-line-client/api/todo-api-client"
)

const MAX_NUMBER_OF_TODOS = 30
const DEFAULT_NUMBER_OF_TODOS = 20

type todosHandler struct {
	ids   []int
	todos []api.ToDoResponse
}

func NewEvenTODOs(from, numberOfIds int) *todosHandler {
	ids := make([]int, 0, MAX_NUMBER_OF_TODOS)
	if from%2 != 0 {
		from++
	}
	for i := 0; i < numberOfIds; i++ {
		ids = append(ids, from+i*2)
	}
	return &todosHandler{
		ids: ids,
	}
}

func NewOddTODOs(from, numberOfIds int) *todosHandler {
	ids := make([]int, 0, MAX_NUMBER_OF_TODOS)
	if from%2 == 0 {
		from++
	}
	for i := 0; i < numberOfIds; i++ {
		ids = append(ids, from+i*2)
	}
	return &todosHandler{
		ids: ids,
	}
}

func NewCustomTODOs(idsRange string) *todosHandler {
	idStrs := strings.Split(idsRange, ",")
	ids := make([]int, 0, MAX_NUMBER_OF_TODOS)

	for _, s := range idStrs {
		id, err := strconv.Atoi(s)
		if err == nil {
			if id >= 1 && id <= 100 {
				ids = append(ids, id)
			}
		}
		if len(ids) == MAX_NUMBER_OF_TODOS {
			break
		}
	}
	return &todosHandler{
		ids: ids,
	}
}

func (t *todosHandler) GetTodos() {
	client := api.NewToDoApiClient()
	cases := make([]reflect.SelectCase, len(t.ids))
	for i, id := range t.ids {
		c := client.GetTODO(id)
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		}
	}
	t.todos = make([]api.ToDoResponse, 0, len(cases))
	for len(cases) > 0 {
		i, v, ok := reflect.Select(cases)
		if !ok {
			cases = append(cases[:i], cases[i+1:]...)
			continue
		}
		t.todos = append(t.todos, v.Interface().(api.ToDoResponse))
	}
}

func (t *todosHandler) ToDosReport() {
	for _, todo := range t.todos {
		if todo.ErrorMessage == "" {
			todo.ToDo.Print()
		} else {
			todo.Error()
		}
		fmt.Println("---------------------------------------")
	}
	fmt.Printf("Total: %d\n", len(t.todos))
}
