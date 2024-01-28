package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
    "net/http"
)

func main() {
    resp, err := http.Get("http://example.com/")
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    fmt.Println(resp.Body)

    cmd := exec.Command("notify-send", "hello")
    cmd.Run()

    for {
        fmt.Println("Hello, world!")
        time.Sleep(1 * time.Second)
    }

    if len(os.Args) % 2 != 1 {
        fmt.Println("Please provide an even number of arguments")
        os.Exit(1)
    }

    for _, arg := range os.Args[1:] {
        fmt.Println(arg)
    }
}
