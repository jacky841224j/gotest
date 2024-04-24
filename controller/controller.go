package controller

import (
	"encoding/json"
	"fmt"
	"gotest/services"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var TodoList []Todo

type Todo struct {
	Id   int64
	Item string
}

type ApiResponse struct {
	ResultCode    string
	ResultMessage interface{}
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(io.LimitReader(r.Body, 1024)) //io.LimitReader限制大小
	if err != nil {
		fmt.Println(err)
	}
	var addTodo Todo
	_ = json.Unmarshal(body, &addTodo) //轉為json
	TodoList = append(TodoList, addTodo)
	defer r.Body.Close()

	response := ApiResponse{"200", TodoList}
	services.ResponseWithJson(w, http.StatusOK, response) //回
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"]
	var targetTodo Todo

	for _, Todo := range TodoList {
		intQueryId, _ := strconv.ParseInt(queryId, 10, 60)
		if Todo.Id == intQueryId {
			targetTodo = Todo
		}
	}

	response := ApiResponse{"200", targetTodo}
	services.ResponseWithJson(w, http.StatusOK, response)
}

func Test(w http.ResponseWriter, r *http.Request) {

	response := ApiResponse{"200", "hello world"}
	services.ResponseWithJson(w, http.StatusOK, response)
}
