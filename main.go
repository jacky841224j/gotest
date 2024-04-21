package main

import (
	"fmt"
	"net/http"

	"gotest/dto"

	routes "gotest/router"
)

func main() {
	person := dto.Person{Name: "John", Age: 30}

	var x int = 2
	var xptr *int = &x
	add(xptr)
	fmt.Println(x)
	fmt.Print(person)

	router := routes.NewRouter()
	http.ListenAndServe(":3000", router)

}

func add(xptr *int) {

	fmt.Println(*xptr)
}
