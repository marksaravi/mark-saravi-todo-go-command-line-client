package api

import "fmt"

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

func NewToDoApiClient(baseUrl string) *todoApiClient {
	return &todoApiClient{
		baseUrl: baseUrl,
	}
}

func (client *todoApiClient) GetTODOs(ids ...int) []ToDoResponse {
	for _, id := range ids {
		fmt.Println(id)
	}
	return []ToDoResponse{}
}
