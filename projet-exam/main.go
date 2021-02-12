package main

import (
	"bytes"
	"encoding/binary"
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
		if element.Done == false {
			chaine += fmt.Sprintf("ID: %d, task: \"%s\"\n", index, element.Description)
		}
	}
	_, err := rw.Write([]byte(chaine)) //call to rw.WriteHeader(http.StatusOK) made within Write()
	//println(chaine) test
	if err != nil {
		fmt.Println(err)
	}
}

func add(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("entering add")
	if req.Method == http.MethodPost {
		body, err := ioutil.ReadAll(req.Body)
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
			Done:        false,
		})
		rw.WriteHeader(http.StatusOK)
		return
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("bad request"))
	}
}

func done(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("entering done")
	switch req.Method {
	case http.MethodGet:
		var chaine = ""
		for index, element := range task {
			if element.Done == true {
				chaine += fmt.Sprintf("ID: %d, task: \"%s\"\n", index, element.Description)
			}
		}
		_, err := rw.Write([]byte(chaine)) //call to rw.WriteHeader(http.StatusOK) made within Write()
		//println(chaine) test
		if err != nil {
			fmt.Println(err)
		}
	case http.MethodPost:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("Error reading body: %v", err)
			http.Error(
				rw,
				"can't read body", http.StatusBadRequest,
			)
			return
		}
		var nombre int
		binary.Read(bytes.NewReader(body), binary.BigEndian, &nombre)
		if nombre < len(task) {
			task[nombre].Done = true
		} else {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("bad request"))
		}
	default:
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("bad request"))
	}
}

func main() {
	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
	fmt.Println("launching server")

	http.ListenAndServe(":8080", nil) //nil -> mux is used
}
