package todos

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	api "gitgub.com/marksaravi/mark-saravi-todo-go-command-line-client/api/todo-api-client"
)

const MAX_NUMBER_OF_TODOS = 30
const DEFAULT_NUMBER_OF_TODOS = 20

type todosHandler struct {
	ids []int
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

func (t *todosHandler) ToDosReport() {
	fmt.Println(os.Getenv("BASE_URL"))
	client := api.NewToDoApiClient()
	for res := range client.MockGetTODO(t.ids[0]) {
		fmt.Println(res.HTTPStatusCode)
		fmt.Println(res.ErrorMessage)
		res.ToDo.Print()
		fmt.Print("__________________________________________________\n")
	}
}
