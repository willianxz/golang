package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Create("data.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()

    fmt.Fprintln(file, "nwsHexKey = a8773564ebc8f7abdcaac6bd2137dd07")
    fmt.Fprintln(file, "appHexKey = 1e745697990164f22531bf11e4614ad1")
    fmt.Fprintln(file, "devHexAddr = 018a6355")
    fmt.Fprintln(file, "broker = tcp://localhost:1884")
    fmt.Fprintln(file, "username = ")
    fmt.Fprintln(file, "password = ")
}
