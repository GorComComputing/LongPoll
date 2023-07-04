package main

import (
    "fmt"
    "net/http"
    "log"
    "io"
    "io/ioutil"
    "strconv"
)

var messages chan string = make(chan string, 100)

var counter = 0

func PushHandler(w http.ResponseWriter, req *http.Request) {
	
	fmt.Println(req.Body)

    _, err := ioutil.ReadAll(req.Body)

    if err != nil {
        w.WriteHeader(400)
    }
    counter += 1
    messages <- strconv.Itoa(counter)
    
    fmt.Println(counter)
    
}


func PollResponse(w http.ResponseWriter, req *http.Request) {

fmt.Println("PollResp")

    io.WriteString(w, <-messages)
}

func main() {

    fmt.Println("http://localhost:8005/")
    
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.HandleFunc("/poll", PollResponse)
    http.HandleFunc("/push", PushHandler)
    err := http.ListenAndServe(":8005", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
