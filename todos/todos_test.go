package todos

import (
	"testing"
	"time"

	api "gitgub.com/marksaravi/mark-saravi-todo-go-command-line-client/api/todo-api-client"
)

func TestWaitTodo(t *testing.T) {
	channel1 := make(chan api.ToDoResponse)
	channel2 := make(chan api.ToDoResponse)
	handler := todosHandler{
		ids: []int{1, 2},
		todoChannels: []<-chan api.ToDoResponse{
			channel1,
			channel2,
		},
	}
	go func() {
		channel1 <- api.ToDoResponse{
			Id:             1,
			HTTPStatusCode: 200,
			ToDo: api.ToDo{
				Id:        1,
				UserId:    10,
				Title:     "Task 1",
				Completed: true,
			},
		}
		channel1 <- api.ToDoResponse{
			Id:             2,
			HTTPStatusCode: 200,
			ToDo: api.ToDo{
				Id:        2,
				UserId:    20,
				Title:     "Task 2",
				Completed: false,
			},
		}
		time.Sleep(100 * time.Millisecond)
		close(channel1)
		close(channel2)
	}()
	handler.WaitTodos()
	if len(handler.todos) != 2 {
		t.Errorf("Channels not received")
	}
}
