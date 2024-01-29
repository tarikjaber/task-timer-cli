
package main

import (
    "os"
    "log"
    "fmt"
)

func writeString(s string) {
    home := os.Getenv("HOME")
    stateFile := fmt.Sprintf("%s/.config/waybar/repeat_state", home)

    file, err := os.Create(stateFile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    _, err = file.WriteString(s)
    if err != nil {
        log.Fatal(err)
    }
}

