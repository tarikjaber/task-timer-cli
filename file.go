
package main

import (
    "os"
    "log"
)

func writeString(s string) {
    stateFile := "/home/tarik/.config/waybar/repeat_state"

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

