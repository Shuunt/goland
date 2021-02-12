package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var task []Task

type Task struct {
	Description string
	Done        bool
}

func list(rw http.ResponseWriter, _ *http.Request) {
	fmt.Println("entering list")
	var chaine = ""
	for index, element := range task {
		if element.Done {
			chaine += fmt.Sprintf("ID: %d, task: \"%s\"\n", index, element.Description)
		}
	}
	_, err := rw.Write([]byte(chaine)) //call to rw.WriteHeader(http.StatusOK) made within Write()
	//println(chaine) test
	if err != nil {
		fmt.Println(err)
	}
}

func add(rw http.ResponseWriter, la *http.Request) {
	fmt.Println("entering add")
	if la.Method == http.MethodPost {
		body, err := ioutil.ReadAll(la.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v", err)
			http.Error(
				rw,
				"can't read body", http.StatusBadRequest,
			)
			return
		}
		task = append(task, Task{
			Description: string(body),
			Done:        false, //welp, no wonder it won't show up if you want to list it after
		})
		rw.WriteHeader(http.StatusOK)
		return
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("bad request"))
	}
}

func done(rw http.ResponseWriter, _ *http.Request) {
	fmt.Println("entering done")
}

func main() {

	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
	fmt.Println("launching server")

	http.ListenAndServe(":8080", nil) //nil -> mux is used
}
