package api

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type ToDo struct {
	Id        int
	UserId    int
	Completed bool
	Title     string
}

type ToDoResponse struct {
	HTTPStatusCode int
	ErrorMessage   string
	ToDo           ToDo
}

type todoApiClient struct {
	baseUrl string
}

func NewToDoApiClient() *todoApiClient {
	return &todoApiClient{
		baseUrl: os.Getenv("TODO_BASE_URL"),
	}
}

func (client *todoApiClient) GetTODOs(ids ...int) []ToDoResponse {
	for _, id := range ids {
		fmt.Println(id)
	}
	return []ToDoResponse{}
}

func (client *todoApiClient) GetTODO(id int) <-chan ToDoResponse {
	out := make(chan ToDoResponse, 1)
	go func(out chan ToDoResponse) {
		defer close(out)

		url := fmt.Sprintf("%s/%d", client.baseUrl, id)
		req, _ := http.NewRequest(http.MethodGet, url, nil)
		req.Header.Set("Content-Type", "application/json")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			out <- ToDoResponse{
				HTTPStatusCode: res.StatusCode,
				ErrorMessage:   err.Error(),
			}
			return
		}
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			out <- ToDoResponse{
				HTTPStatusCode: res.StatusCode,
				ErrorMessage:   err.Error(),
			}
			return
		}
		var todo ToDo
		err = json.Unmarshal(resBody, &todo)
		if err != nil {
			out <- ToDoResponse{
				HTTPStatusCode: res.StatusCode,
				ErrorMessage:   err.Error(),
			}
		}
		out <- ToDoResponse{
			HTTPStatusCode: res.StatusCode,
			ErrorMessage:   "",
			ToDo:           todo,
		}
	}(out)
	return out
}

func (client *todoApiClient) MockGetTODO(id int) <-chan ToDoResponse {
	out := make(chan ToDoResponse, 1)
	go func(out chan ToDoResponse) {
		defer close(out)

		time.Sleep(time.Millisecond*50 + time.Millisecond*time.Duration(rand.Intn(50)))
		out <- ToDoResponse{
			HTTPStatusCode: http.StatusOK,
			ErrorMessage:   "",
			ToDo: ToDo{
				Id:        id,
				UserId:    id,
				Title:     fmt.Sprintf("Task #%d", id),
				Completed: rand.Intn(100)%2 == 0,
			},
		}

	}(out)
	return out
}

func (t ToDo) Print() {
	blue := color.New(color.FgBlue)
	yellow := color.New(color.FgYellow)

	completed := "Yes"
	if !t.Completed {
		completed = "No"
	}
	blue.Print("Id       ")
	fmt.Print(": ")
	yellow.Println(t.Id)
	blue.Print("Title    ")
	fmt.Print(": ")
	yellow.Println(t.Title)
	blue.Print("Completed")
	fmt.Print(": ")
	yellow.Println(completed)
}
