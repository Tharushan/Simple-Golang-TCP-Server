package main

import (
  "io"
  "log"
  "net"
  "os"
  "time"
)

func main() {
  l, err := net.Listen("tcp", "localhost:5432")
  if err != nil {
      log.Fatal(err)
  }
  defer l.Close()
  for {
    conn, err := l.Accept()
    if err != nil {
      log.Fatal(err)
    }
    go func(c net.Conn) {
      go func() {
        time.Sleep(time.Second * 1)
        c.Write([]byte("Hello, message received."))
      }()
      defer c.Close()
      io.Copy(os.Stdout, c)
    }(conn)
  }
}
