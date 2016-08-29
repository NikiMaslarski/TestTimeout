package main

import (
  "net"
  "log"
  "time"
  "fmt"
  "bufio"
)

func main() {
  req := `
  GET / HTTP/1.1
  Host: localhost:8080
  User-Agent: curl/7.47.0
  Accept: *//*
`


  conn, err := net.Dial("tcp", "127.0.0.1:8080")
  handleErr(err)
  start := time.Now()
  n, err := conn.Read([]byte(" "))
  elapsed := time.Since(start)
  log.Println("Read timeout test")
  log.Println(fmt.Sprintf("Err is %s, elapsed is: %s, we have read %d bytes ", err, elapsed, n))

  start = time.Now()
  buf := make([]byte, 255)
  n, err = conn.Write([]byte(req))
  log.Println(fmt.Sprintf("Err is %s, we have write %d bytes ", err, n))

  n, err = conn.Read(buf)
  fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
  resp, err := bufio.NewReader(conn).ReadString('\n')
  log.Println(fmt.Sprintf("resp: %s, err: %s", resp, err))

  elapsed = time.Since(start)
  log.Println("Write timeout test")
  log.Println(fmt.Sprintf("Err is: %s", err))
  log.Println(fmt.Sprintf("We have read : %d bytes", n))
  log.Println(fmt.Sprintf("Buffer is:\n%s ", buf))
  log.Println(fmt.Sprintf("elapsed is: %s", elapsed))




}

func handleErr(err error){
  if err != nil {
    log.Fatal(err)
  }
}
