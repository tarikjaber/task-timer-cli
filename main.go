package main

import (
    "fmt"
    "os"
    "os/exec"
    "strconv"
    "time"
)

func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func main() {
    if len(os.Args) % 2 != 1 {
        fmt.Println("Please provide an even number of arguments")
        os.Exit(1)
    }

    for i := 1; i < len(os.Args); i += 2 {
        clearScreen()
        timeUnit := "minutes"
        if os.Args[i + 1] == "1" {
            timeUnit = "minute"
        }
        notify := exec.Command("notify-send", fmt.Sprintf("%s for %s %s", os.Args[i], os.Args[i + 1], timeUnit))
        notify.Run()
        fmt.Println(os.Args[i])
        minutes, err := strconv.Atoi(os.Args[i + 1])
        if err != nil {
            minutes = 5
        }
        totalSeconds := minutes * 60
        for totalSeconds > 0 {
            fmt.Printf("\r%02d:%02d", totalSeconds / 60, totalSeconds % 60)
            totalSeconds -= 1
            time.Sleep(1 * time.Second)
        }
    }
    fmt.Println("\nDone!")
}
