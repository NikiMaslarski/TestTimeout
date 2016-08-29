package main

import (
  "net/http"
  "time"
  "fmt"
  "log"
  "os"
)

func main() {
  myMux := http.NewServeMux()
  myMux.HandleFunc("/", someFunc)


  server := http.Server{
    Addr: ":8080",
    ReadTimeout: time.Duration(10) * time.Second,
    WriteTimeout: time.Duration(5) * time.Second,
    Handler: myMux,
  }


  err := server.ListenAndServe()
  if err != nil {
    fmt.Println("Error is %v", err)
    log.Fatal(err)
  }
}

func someFunc(w http.ResponseWriter, req *http.Request) {
  var pid, ppid int
  pid = os.Getpid()
  ppid = os.Getppid()
  w.Write([]byte(fmt.Sprintf("pid is: %d \nppid is: %d", pid, ppid)))
}
